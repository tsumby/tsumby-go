package tsumby

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testBasePath = "http://localhost:8000/"
	testSecret   = "secret"
)

func TestRequestCreate(t *testing.T) {
	client := New(testSecret)
	client.Debug = true
	client.BaseURL = testBasePath

	img, err := client.Create(context.Background(), Params{
		Image: "https://raw.githubusercontent.com/cshum/imagor/master/testdata/dancing-banana.gif",
		Width: 500,
		Filters: Filters{
			Filter{
				Name: "quality",
				Args: "50",
			},
			Filter{
				Name: "fill",
				Args: "yellow",
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, img)

	fmt.Println(img.Type)
}
