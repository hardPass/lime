package sublime

import (
	"fmt"
	"lime/3rdparty/libs/gopy/lib"
	"lime/backend"
	"lime/backend/primitives"
)

var (
	_ = backend.View{}
	_ = primitives.Region{}
	_ = fmt.Errorf
)

func (o *Window) Py_active_view() (py.Object, error) {
	ret0 := backend.GetEditor().Frontend().ActiveView(o.data)
	var err error
	var pyret0 py.Object

	pyret0, err = _viewClass.Alloc(1)
	if err != nil {
	} else if v2, ok := pyret0.(*View); !ok {
		return nil, fmt.Errorf("Unable to convert return value to the right type?!: %s", pyret0.Type())
	} else {
		v2.data = ret0
	}
	if err != nil {
		return nil, err
	}
	return pyret0, err
}

func (o *Window) Py_run_command(tu *py.Tuple) (py.Object, error) {
	var (
		arg1 string
		arg2 backend.Args
	)
	if v, err := tu.GetItem(0); err != nil {
		return nil, err
	} else {
		if v2, ok := v.(*py.String); !ok {
			return nil, fmt.Errorf("Expected type *py.String for backend.Window.RunCommand() arg1, not %s", v.Type())
		} else {
			arg1 = v2.String()
		}
	}
	arg2 = make(backend.Args)
	if v, err := tu.GetItem(1); err == nil {
		if v2, ok := v.(*py.Dict); !ok {
			return nil, fmt.Errorf("Expected type *py.Dict for backend.Window.RunCommand() arg2, not %s", v.Type())
		} else {
			if v, err := fromPython(v2); err != nil {
				return nil, err
			} else {
				arg2 = v.(backend.Args)
			}
		}
	}
	backend.GetEditor().CommandHandler().RunWindowCommand(o.data, arg1, arg2)
	py.None.Incref()
	return py.None, nil
}