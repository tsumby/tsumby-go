package tsumby

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testBasePath = "http://localhost:8000/"
	testSecret   = "secret"
)

func TestRequestCreate(t *testing.T) {
	client := New(testSecret)
	client.BaseURL = testBasePath

	img, err := client.Create(context.Background(), Params{
		Image: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		Width: 50,
	})
	require.NoError(t, err)
	require.NotNil(t, img)
}
