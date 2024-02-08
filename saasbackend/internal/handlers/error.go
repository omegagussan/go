package handlers

type NotFoundError struct {
    Msg  string
}

func (e NotFoundError) Error() string {
    return e.Msg
}

type BadDataError struct {
    Msg  string
}

func (e BadDataError) Error() string {
    return e.Msg
}