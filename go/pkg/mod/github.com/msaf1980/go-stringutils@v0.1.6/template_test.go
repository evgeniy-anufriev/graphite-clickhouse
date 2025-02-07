package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate_NewTemplate(t *testing.T) {
	tests := []struct {
		format   string
		template Template
		wantErr  bool
	}{
		{format: "", template: Template{"", ""}, wantErr: false},
		{format: "test", template: Template{"test", "test"}, wantErr: false},
		{
			format: "test %{query} %{log.status}",
			template: Template{
				"test %{query} %{log.status}",
				"test ", templateParam{"query", "query"}, " ", templateParam{"log.status", "log", "status"},
			},
			wantErr: false,
		},
		{
			format: "%{query}%{log.status} ",
			template: Template{
				"%{query}%{log.status} ",
				templateParam{"query", "query"}, templateParam{"log.status", "log", "status"}, " ",
			},
			wantErr: false,
		},
		{format: "%query}", template: nil, wantErr: true},
		{format: "%{query", template: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			tpl, err := NewTemplate(tt.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equalf(t, tt.template, tpl, "NewTemplate() results mismatch")
		})
	}
}

func TestTemplate_Execute(t *testing.T) {
	params := map[string]interface{}{
		"query":  "URL",
		"status": "success",
		"time":   "now",
		"one": map[string]interface{}{
			"two_one": "2_1", "two_two": "2_2",
			"three": map[string]interface{}{"three_four": "3_4"},
		},
		"int32":   int32(-1),
		"uint32":  uint32(1),
		"int64":   int64(-2),
		"uint64":  uint64(2),
		"int16":   int16(-1),
		"uint16":  uint16(1),
		"int8":    int8(-1),
		"uint8":   uint8(1),
		"int":     -1,
		"uint":    1,
		"float32": 1.3,
		"float64": 2.4,
	}
	tests := []struct {
		format  string
		want    string
		wantErr bool
	}{
		// one element
		{format: "", want: "", wantErr: false},
		{format: "test", want: "test", wantErr: false},
		{format: "%{query}", want: "URL", wantErr: false},
		{format: "%{not_exist}", want: "", wantErr: true},
		//several elements
		{format: "%{query} %{status}", want: "URL success", wantErr: false},
		{format: "%{query} %{status} at %{time}", want: "URL success at now", wantErr: false},
		// deep params
		{format: "%{query} %{one.two_one} %{one.three.three_four}", want: "URL 2_1 3_4", wantErr: false},
		{format: "%{query} %{one.two_one} %{one.three.three_four.2}", want: "", wantErr: true},
		{format: "%{query} %{one.two_one} %{one.three}", want: "", wantErr: true},
		// non-string data types
		{format: "%{int32} %{uint32} %{int64} %{uint64} %{float32} %{float64}", want: "-1 1 -2 2 1.3 2.4", wantErr: false},
		{format: "%{int16} %{uint16}", want: "-1 1", wantErr: false},
		{format: "%{int8} %{uint8}", want: "-1 1", wantErr: false},
		{format: "%{int} %{uint}", want: "-1 1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			tpl, err := NewTemplate(tt.format)
			if err != nil {
				t.Fatalf("NewTemplate() error = %v", err)
			}
			got, err := tpl.Execute(params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Template.Execute() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func TestTemplate_ExecutePartial(t *testing.T) {
	params := map[string]interface{}{
		"query":  "URL",
		"status": "success",
		"time":   "now",
		"one": map[string]interface{}{
			"two_one": "2_1", "two_two": "2_2",
			"three": map[string]interface{}{"three_four": "3_4"},
		},
		"int32":   int32(-1),
		"uint32":  uint32(1),
		"int64":   int64(-2),
		"uint64":  uint64(2),
		"int16":   int16(-1),
		"uint16":  uint16(1),
		"int8":    int8(-1),
		"uint8":   uint8(1),
		"int":     -1,
		"uint":    1,
		"float32": 1.3,
		"float64": 2.4,
	}
	tests := []struct {
		format      string
		want        string
		wantPartial bool
	}{
		// one element
		{format: "", want: "", wantPartial: false},
		{format: "test", want: "test", wantPartial: false},
		{format: "%{query}", want: "URL", wantPartial: false},
		{format: "%{not_exist}", want: "%{not_exist}", wantPartial: true},
		//several elements
		{format: "%{query} %{status}", want: "URL success", wantPartial: false},
		{format: "%{query} %{status} at %{time}", want: "URL success at now", wantPartial: false},
		// deep params
		{format: "%{query} %{one.two_one} %{one.three.three_four}", want: "URL 2_1 3_4", wantPartial: false},
		{format: "%{query} %{one.two_one} %{one.three.three_four.2}", want: "URL 2_1 %{one.three.three_four.2}", wantPartial: true},
		{format: "%{query} %{one.two_one} %{one.three}", want: "URL 2_1 %{one.three}", wantPartial: true},
		// non-string data types
		{format: "%{int32} %{uint32} %{int64} %{uint64} %{float32} %{float64}", want: "-1 1 -2 2 1.3 2.4", wantPartial: false},
		{format: "%{int16} %{uint16}", want: "-1 1", wantPartial: false},
		{format: "%{int8} %{uint8}", want: "-1 1", wantPartial: false},
		{format: "%{int} %{uint}", want: "-1 1", wantPartial: false},
	}
	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			tpl, err := NewTemplate(tt.format)
			if err != nil {
				t.Fatalf("NewTemplate() error = %v", err)
			}
			got, part := tpl.ExecutePartial(params)
			if part != tt.wantPartial {
				t.Errorf("Template.ExecutePartial() partial = %v, want %v", err, tt.wantPartial)
				return
			}
			if got != tt.want {
				t.Errorf("Template.Execute() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Benchmark_TemplateExecute(b *testing.B) {
	params := map[string]interface{}{
		"query":  "URL",
		"status": "success",
		"time":   "now",
		"one": map[string]interface{}{
			"two_one": "2_1", "two_two": "2_2",
			"three": map[string]interface{}{"three_four": "3_4"},
		},
	}
	format := "%{query} %{one.two_one} %{one.three.three_four}"
	tpl, err := NewTemplate(format)
	if err != nil {
		b.Fatalf("NewTemplate() error = %v", err)
	}
	for i := 0; i < b.N; i++ {
		got, err := tpl.Execute(params)
		_ = got
		if err != nil {
			b.Fatalf("Template.Execute() error = %v", err)
			return
		}
	}
}

func Benchmark_TemplateInitAndExecute(b *testing.B) {
	params := map[string]interface{}{
		"query":  "URL",
		"status": "success",
		"time":   "now",
		"one": map[string]interface{}{
			"two_one": "2_1", "two_two": "2_2",
			"three": map[string]interface{}{"three_four": "3_4"},
		},
	}
	format := "%{query} %{one.two_one} %{one.three.three_four}"
	for i := 0; i < b.N; i++ {
		tpl, err := NewTemplate(format)
		if err != nil {
			b.Fatalf("NewTemplate() error = %v", err)
		}
		got, err := tpl.Execute(params)
		_ = got
		if err != nil {
			b.Fatalf("Template.Execute() error = %v", err)
			return
		}
	}
}
