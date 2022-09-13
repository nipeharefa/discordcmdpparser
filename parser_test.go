package discordcmdpparser

import (
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		want    ParserResult
		wantErr bool
	}{
		{
			args: args{
				content: "!ping",
			},
			want: ParserResult{
				Command: "!ping",
			},
			name: "test ping",
		},
		{
			args: args{
				content: "!call gowi",
			},
			want: ParserResult{
				Command: "!call",
				Arg:     "gowi",
			},
			name: "call with args",
		},
		{
			args: args{
				content: "call gowi",
			},
			want:    ParserResult{},
			wantErr: true,
			name:    "parsing invalid command",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parser(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invalid %v want %v", []byte(got.Arg), []byte(tt.want.Arg))
				t.Errorf("Parser() = %v, want %v", got, tt.want)
			}
		})
	}
}
