package types

var Types = map[string]string{
	"int":           "reflect.TypeOf(int(0))",
	"unsigned int":  "reflect.TypeOf(uint32(0))",
	"char":          "reflect.TypeOf(int8(0))",
	"unsigned char": "reflect.TypeOf(uint8(0))",
	"size_t":        "reflect.TypeOf(uint64(0))",
	"char *":        `reflect.TypeOf("")`,
	"long":          "reflect.TypeOf(int32(0))",
	"unsigned long": "reflect.TypeOf(uint64(0))",
	"off_t":         "reflect.TypeOf(uint32(0))",
	"loff_t":        "reflect.TypeOf(uint64(0))",
	"key_t":         "reflect.TypeOf(int32(0))",
	"pid_t":         "reflect.TypeOf(int32(0))",
	"mode_t":        "reflect.TypeOf(uint16(0))",
	"uid_t":         "reflect.TypeOf(uint32(0))",
	"gid_t":         "reflect.TypeOf(uint32(0))",
	"dev_t":         "reflect.TypeOf(uint32(0))",
	"qid_t":         "reflect.TypeOf(uint32(0))",
	"time_t":        "reflect.TypeOf(uint32(0))",
	"clockid_t":     "reflect.TypeOf(uint32(0))",
	"timer_t":       "reflect.TypeOf(int32(0))",
	"mqd_t":         "reflect.TypeOf(int32(0))",
	"key_serial_t":  "reflect.TypeOf(int32(0))",
	"aio_context_t": "reflect.TypeOf(uint32(0))",
	"void":          "reflect.TypeOf(uint8(0))",
	"u32":           "reflect.TypeOf(uint32(0))",
	"u64":           "reflect.TypeOf(uint64(0))",
	"__s32":         "reflect.TypeOf(int32(0))",
	"__u64":         "reflect.TypeOf(uint64(0))",
	"void *":        `reflect.TypeOf(uintptr(0))`,

	"sigset_t":          "reflect.TypeOf(struct{}{})",
	"fd_set":            "reflect.TypeOf(struct{}{})",
	"cap_user_header_t": "reflect.TypeOf(struct{}{})",
	"cap_user_data_t":   "reflect.TypeOf(struct{}{})",
	"siginfo_t":         "reflect.TypeOf(struct{}{})",
	"stack_t":           "reflect.TypeOf(struct{}{})",
}
