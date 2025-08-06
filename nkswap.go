package libnk

func NkSwap(a *int, b *int) {
	*a, *b = *b, *a
}
