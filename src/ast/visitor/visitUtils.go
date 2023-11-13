package visitor

func typeCastToPtr[T any](i any) *T {
	if i == nil {
		return nil
	}
	return i.(*T)
}

func typeCastToArr[T any](v any) []T {
	if v == nil {
		return nil
	}
	return v.([]T)
}
