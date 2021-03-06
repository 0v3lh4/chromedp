package chromedp

import (
	"context"
	"errors"
	"time"

	"github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/dom"
	"github.com/knq/chromedp/cdp/input"
)

// Error types.
var (
	ErrInvalidDimensions = errors.New("invalid box dimensions")
)

// MouseAction is a mouse action.
func MouseAction(typ input.MouseType, x, y int64, opts ...MouseOption) Action {
	f := input.DispatchMouseEvent(typ, x, y)
	for _, o := range opts {
		f = o(f)
	}

	return f
}

// MouseClickXY sends a left mouse button click at the X, Y location.
func MouseClickXY(x, y int64, opts ...MouseOption) Action {
	return Tasks{
		MouseAction(input.MousePressed, x, y, append(opts, Button(input.ButtonLeft), ClickCount(1))...),
		MouseAction(input.MouseReleased, x, y, append(opts, Button(input.ButtonLeft), ClickCount(1))...),
	}
}

// MouseActionNode dispatches a mouse event at the center of a specified node.
func MouseActionNode(n *cdp.Node, opts ...MouseOption) Action {
	return ActionFunc(func(ctxt context.Context, h cdp.FrameHandler) error {
		box, err := dom.GetBoxModel(n.NodeID).Do(ctxt, h)
		if err != nil {
			return err
		}

		c := len(box.Content)
		if c%2 != 0 {
			return ErrInvalidDimensions
		}

		var x, y int64
		for i := 0; i < c; i += 2 {
			x += int64(box.Content[i])
			y += int64(box.Content[i+1])
		}

		return MouseClickXY(x/int64(c/2), y/int64(c/2), opts...).Do(ctxt, h)
	})
}

// MouseOption is a mouse action option.
type MouseOption func(*input.DispatchMouseEventParams) *input.DispatchMouseEventParams

// Button is a mouse action option to set the button to click.
func Button(button input.ButtonType) MouseOption {
	return func(p *input.DispatchMouseEventParams) *input.DispatchMouseEventParams {
		return p.WithButton(button)
	}
}

// ButtonString is a mouse action option to set the button to click as a
// string.
func ButtonString(btn string) MouseOption {
	return Button(input.ButtonType(btn))
}

// ButtonModifiers is a mouse action option to add additional modifiers for the
// button.
func ButtonModifiers(modifiers ...input.Modifier) MouseOption {
	return func(p *input.DispatchMouseEventParams) *input.DispatchMouseEventParams {
		for _, m := range modifiers {
			p.Modifiers |= m
		}
		return p
	}
}

// ClickCount is a mouse action option to set the click count.
func ClickCount(n int) MouseOption {
	return func(p *input.DispatchMouseEventParams) *input.DispatchMouseEventParams {
		return p.WithClickCount(int64(n))
	}
}

// KeyAction contains information about a key action.
type KeyAction struct {
	v    string
	opts []KeyOption
}

// KeyCode are known system key codes.
type KeyCode string

// KeyCode values.
const (
	KeyCodeBackspace = "\b"
	KeyCodeTab       = "\t"
	KeyCodeCR        = "\r"
	KeyCodeLF        = "\n"
	KeyCodeLeft      = "\x25"
	KeyCodeUp        = "\x26"
	KeyCodeRight     = "\x27"
	KeyCodeDown      = "\x28"
)

const (
	keyRuneCR = '\r'
)

// keyCodeNames is the map of key code values to their respective named
// identifiers.
var keyCodeNames = map[KeyCode]string{
	KeyCodeBackspace: "Backspace",
	KeyCodeTab:       "Tab",
	KeyCodeCR:        "Enter",
	KeyCodeLF:        "Enter",
	KeyCodeLeft:      "Left",
	KeyCodeUp:        "Up",
	KeyCodeRight:     "Right",
	KeyCodeDown:      "Down",
}

// Do satisfies Action interface.
func (ka *KeyAction) Do(ctxt context.Context, h cdp.FrameHandler) error {
	var err error

	// apply opts
	sysP := input.DispatchKeyEvent(input.KeyRawDown)
	keyP := input.DispatchKeyEvent(input.KeyChar)
	for _, o := range ka.opts {
		sysP = o(sysP)
		keyP = o(keyP)
	}

	for _, r := range ka.v {
		s := string(r)
		keyS := KeyCode(r)
		if n, ok := keyCodeNames[keyS]; ok {
			kc := int64(r)
			if keyS == KeyCodeLF {
				s = string(keyRuneCR)
				kc = int64(keyRuneCR)
			}

			err = sysP.WithKey(n).
				WithNativeVirtualKeyCode(kc).
				WithWindowsVirtualKeyCode(kc).
				WithKeyIdentifier(s).
				WithIsSystemKey(true).
				Do(ctxt, h)
			if err != nil {
				return err
			}
		}

		err = keyP.WithText(s).Do(ctxt, h)
		if err != nil {
			return err
		}

		// FIXME
		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

// KeyActionNode dispatches a key event on a node.
func KeyActionNode(n *cdp.Node, v string, opts ...KeyOption) Action {
	return Tasks{
		dom.Focus(n.NodeID),
		MouseActionNode(n),
		&KeyAction{v, opts},
	}
}

// KeyOption is a key action option.
type KeyOption func(*input.DispatchKeyEventParams) *input.DispatchKeyEventParams

// KeyModifiers is a key action option to add additional modifiers on the key
// press.
func KeyModifiers(modifiers ...input.Modifier) KeyOption {
	return func(p *input.DispatchKeyEventParams) *input.DispatchKeyEventParams {
		for _, m := range modifiers {
			p.Modifiers |= m
		}
		return p
	}
}
