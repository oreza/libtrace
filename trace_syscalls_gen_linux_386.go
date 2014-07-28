package libtrace

var (
	type_int          = int(0)
	type_uint         = uint(0)
	type_int8         = int8(0)
	type_int16        = int16(0)
	type_int32        = int32(0)
	type_int64        = int64(0)
	type_uint8        = uint8(0)
	type_uint16       = uint16(0)
	type_uint32       = uint32(0)
	type_uint64       = uint64(0)
	type_uintptr      = uintptr(0)
	type_float32      = float32(0)
	type_float64      = float64(0)
	type_stringc      = StringC("")
	type_stringbuffer = StringBuffer("")
	type_buffer       = []byte{}

	type_unknownstruct = struct{}{}
)

var syscalls = []*Signature{

	&Signature{Id: 1, Name: "exit", Args: []Arg{Arg{Name: "error_code", Type: type_int, Const: false}}},
	&Signature{Id: 2, Name: "fork", Args: []Arg{Arg{Name: "regs", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 3, Name: "read", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "buf", Type: type_stringc, Const: false}, Arg{Name: "count", Type: type_uint64, Const: false}}},
	&Signature{Id: 4, Name: "write", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "buf", Type: type_stringc, Const: true}, Arg{Name: "count", Type: type_uint64, Const: false}}},
	&Signature{Id: 5, Name: "open", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: true}, Arg{Name: "flags", Type: type_int, Const: false}, Arg{Name: "mode", Type: type_int, Const: false}}},
	&Signature{Id: 6, Name: "close", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}}},
	&Signature{Id: 7, Name: "waitpid", Args: []Arg{Arg{Name: "pid", Type: type_int32, Const: false}, Arg{Name: "status", Type: &type_uint, Const: false}, Arg{Name: "options", Type: type_int, Const: false}}},
	&Signature{Id: 8, Name: "creat", Args: []Arg{Arg{Name: "pathname", Type: type_stringc, Const: true}, Arg{Name: "mode", Type: type_int, Const: false}}},
	&Signature{Id: 9, Name: "link", Args: []Arg{Arg{Name: "oldname", Type: type_stringc, Const: true}, Arg{Name: "newname", Type: type_stringc, Const: true}}},
	&Signature{Id: 10, Name: "unlink", Args: []Arg{Arg{Name: "pathname", Type: type_stringc, Const: true}}},
	&Signature{Id: 11, Name: "execve", Args: []Arg{Arg{Name: "filename", Type: type_unknownstruct, Const: false}}},
	&Signature{Id: 12, Name: "chdir", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: true}}},
	&Signature{Id: 13, Name: "time", Args: []Arg{Arg{Name: "tloc", Type: &type_int, Const: false}}},
	&Signature{Id: 14, Name: "mknod", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: true}, Arg{Name: "mode", Type: type_int, Const: false}, Arg{Name: "dev", Type: type_uint32, Const: false}}},
	&Signature{Id: 15, Name: "chmod", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: true}, Arg{Name: "mode", Type: type_uint16, Const: false}}},
	&Signature{Id: 16, Name: "lchown", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: true}, Arg{Name: "user", Type: type_uint32, Const: false}, Arg{Name: "group", Type: type_uint32, Const: false}}},
	&unknownSignature, // 16
	&Signature{Id: 18, Name: "stat", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: false}, Arg{Name: "statbuf", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 19, Name: "lseek", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "offset", Type: type_uint32, Const: false}, Arg{Name: "origin", Type: type_uint, Const: false}}},
	&Signature{Id: 20, Name: "getpid", Args: []Arg{}},
	&Signature{Id: 21, Name: "mount", Args: []Arg{Arg{Name: "dev_name", Type: type_stringc, Const: false}, Arg{Name: "dir_name", Type: type_stringc, Const: false}, Arg{Name: "type", Type: type_stringc, Const: false}}},
	&Signature{Id: 22, Name: "oldumount", Args: []Arg{Arg{Name: "", Type: type_stringc, Const: false}}},
	&Signature{Id: 23, Name: "setuid", Args: []Arg{Arg{Name: "uid", Type: type_uint32, Const: false}}},
	&Signature{Id: 24, Name: "getuid", Args: []Arg{}},
	&Signature{Id: 25, Name: "stime", Args: []Arg{Arg{Name: "", Type: &type_int, Const: false}}},
	&Signature{Id: 26, Name: "ptrace", Args: []Arg{Arg{Name: "request", Type: type_int32, Const: false}, Arg{Name: "pid", Type: type_int32, Const: false}, Arg{Name: "addr", Type: type_int32, Const: false}, Arg{Name: "data", Type: type_int32, Const: false}}},
	&Signature{Id: 27, Name: "alarm", Args: []Arg{Arg{Name: "seconds", Type: type_uint, Const: false}}},
	&Signature{Id: 28, Name: "fstat", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "statbuf", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 29, Name: "pause", Args: []Arg{}},
	&Signature{Id: 30, Name: "utime", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: false}, Arg{Name: "times", Type: &type_unknownstruct, Const: false}}},
	&unknownSignature, // 30
	&unknownSignature, // 31
	&Signature{Id: 33, Name: "access", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: true}, Arg{Name: "mode", Type: type_int, Const: false}}},
	&Signature{Id: 34, Name: "nice", Args: []Arg{Arg{Name: "incr", Type: type_int, Const: false}}},
	&unknownSignature, // 34
	&Signature{Id: 36, Name: "sync", Args: []Arg{}},
	&Signature{Id: 37, Name: "kill", Args: []Arg{Arg{Name: "pid", Type: type_int, Const: false}, Arg{Name: "sig", Type: type_int, Const: false}}},
	&Signature{Id: 38, Name: "rename", Args: []Arg{Arg{Name: "oldname", Type: type_stringc, Const: true}, Arg{Name: "newname", Type: type_stringc, Const: true}}},
	&Signature{Id: 39, Name: "mkdir", Args: []Arg{Arg{Name: "pathname", Type: type_stringc, Const: true}, Arg{Name: "mode", Type: type_int, Const: false}}},
	&Signature{Id: 40, Name: "rmdir", Args: []Arg{Arg{Name: "pathname", Type: type_stringc, Const: true}}},
	&Signature{Id: 41, Name: "dup", Args: []Arg{Arg{Name: "fildes", Type: type_uint, Const: false}}},
	&Signature{Id: 42, Name: "pipe", Args: []Arg{Arg{Name: "filedes", Type: &type_uint32, Const: false}}},
	&Signature{Id: 43, Name: "times", Args: []Arg{Arg{Name: "info", Type: &type_unknownstruct, Const: false}}},
	&unknownSignature, // 43
	&Signature{Id: 45, Name: "brk", Args: []Arg{Arg{Name: "brk", Type: type_uint32, Const: false}}},
	&Signature{Id: 46, Name: "setgid", Args: []Arg{Arg{Name: "gid", Type: type_uint32, Const: false}}},
	&Signature{Id: 47, Name: "getgid", Args: []Arg{}},
	&Signature{Id: 48, Name: "signal", Args: []Arg{Arg{Name: "signum", Type: type_int, Const: false}, Arg{Name: "handler", Type: type_unknownstruct, Const: false}}},
	&Signature{Id: 49, Name: "geteuid", Args: []Arg{}},
	&Signature{Id: 50, Name: "getegid", Args: []Arg{}},
	&Signature{Id: 51, Name: "acct", Args: []Arg{Arg{Name: "name", Type: type_stringc, Const: true}}},
	&Signature{Id: 52, Name: "umount", Args: []Arg{Arg{Name: "target", Type: type_stringc, Const: false}, Arg{Name: "flags", Type: type_int, Const: false}}},
	&unknownSignature, // 52
	&Signature{Id: 54, Name: "ioctl", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "cmd", Type: type_uint, Const: false}, Arg{Name: "arg", Type: type_uint32, Const: false}}},
	&Signature{Id: 55, Name: "fcntl", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "cmd", Type: type_uint, Const: false}, Arg{Name: "arg", Type: type_uint32, Const: false}}},
	&unknownSignature, // 55
	&Signature{Id: 57, Name: "setpgid", Args: []Arg{Arg{Name: "pid", Type: type_int32, Const: false}, Arg{Name: "pgid", Type: type_int32, Const: false}}},
	&unknownSignature, // 57
	&Signature{Id: 59, Name: "olduname", Args: []Arg{Arg{Name: "", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 60, Name: "umask", Args: []Arg{Arg{Name: "mask", Type: type_int, Const: false}}},
	&Signature{Id: 61, Name: "chroot", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: true}}},
	&Signature{Id: 62, Name: "ustat", Args: []Arg{Arg{Name: "dev", Type: type_uint32, Const: false}, Arg{Name: "ubuf", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 63, Name: "dup2", Args: []Arg{Arg{Name: "oldfd", Type: type_uint, Const: false}, Arg{Name: "newfd", Type: type_uint, Const: false}}},
	&Signature{Id: 64, Name: "getppid", Args: []Arg{}},
	&Signature{Id: 65, Name: "getpgrp", Args: []Arg{}},
	&Signature{Id: 66, Name: "setsid", Args: []Arg{}},
	&Signature{Id: 67, Name: "sigaction", Args: []Arg{Arg{Name: "signum", Type: type_int, Const: false}, Arg{Name: "act", Type: &type_unknownstruct, Const: true}, Arg{Name: "oldact", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 68, Name: "sgetmask", Args: []Arg{}},
	&Signature{Id: 69, Name: "ssetmask", Args: []Arg{Arg{Name: "signum", Type: type_int, Const: false}}},
	&Signature{Id: 70, Name: "setreuid", Args: []Arg{Arg{Name: "ruid", Type: type_uint32, Const: false}, Arg{Name: "euid", Type: type_uint32, Const: false}}},
	&Signature{Id: 71, Name: "setregid", Args: []Arg{Arg{Name: "rgid", Type: type_uint32, Const: false}, Arg{Name: "egid", Type: type_uint32, Const: false}}},
	&Signature{Id: 72, Name: "sigsuspend", Args: []Arg{Arg{Name: "set", Type: type_int, Const: false}, Arg{Name: "oldset", Type: type_int, Const: false}}},
	&Signature{Id: 73, Name: "sigpending", Args: []Arg{Arg{Name: "set", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 74, Name: "sethostname", Args: []Arg{Arg{Name: "name", Type: type_stringc, Const: false}, Arg{Name: "len", Type: type_int, Const: false}}},
	&Signature{Id: 75, Name: "setrlimit", Args: []Arg{Arg{Name: "resource", Type: type_uint, Const: false}, Arg{Name: "rlim", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 76, Name: "getrlimit", Args: []Arg{Arg{Name: "resource", Type: type_uint, Const: false}, Arg{Name: "rlim", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 77, Name: "getrusage", Args: []Arg{Arg{Name: "who", Type: type_int, Const: false}, Arg{Name: "ru", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 78, Name: "gettimeofday", Args: []Arg{Arg{Name: "tv", Type: &type_unknownstruct, Const: false}, Arg{Name: "tz", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 79, Name: "settimeofday", Args: []Arg{Arg{Name: "tv", Type: &type_unknownstruct, Const: false}, Arg{Name: "tz", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 80, Name: "getgroups", Args: []Arg{Arg{Name: "gidsetsize", Type: type_int, Const: false}, Arg{Name: "grouplist", Type: &type_uint32, Const: false}}},
	&Signature{Id: 81, Name: "setgroups", Args: []Arg{Arg{Name: "gidsetsize", Type: type_int, Const: false}, Arg{Name: "grouplist", Type: &type_uint32, Const: false}}},
	&Signature{Id: 82, Name: "select", Args: []Arg{Arg{Name: "", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 83, Name: "symlink", Args: []Arg{Arg{Name: "oldname", Type: type_stringc, Const: true}, Arg{Name: "newname", Type: type_stringc, Const: true}}},
	&Signature{Id: 84, Name: "lstat", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: false}, Arg{Name: "statbuf", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 85, Name: "readlink", Args: []Arg{Arg{Name: "path", Type: type_stringc, Const: true}, Arg{Name: "buf", Type: type_stringc, Const: false}, Arg{Name: "bufsiz", Type: type_int, Const: false}}},
	&Signature{Id: 86, Name: "uselib", Args: []Arg{Arg{Name: "library", Type: type_stringc, Const: true}}},
	&Signature{Id: 87, Name: "swapon", Args: []Arg{Arg{Name: "specialfile", Type: type_stringc, Const: true}, Arg{Name: "swap_flags", Type: type_int, Const: false}}},
	&Signature{Id: 88, Name: "reboot", Args: []Arg{Arg{Name: "magic1", Type: type_int, Const: false}, Arg{Name: "magic2", Type: type_int, Const: false}, Arg{Name: "cmd", Type: type_int, Const: false}, Arg{Name: "arg", Type: &type_uint8, Const: false}}},
	&Signature{Id: 89, Name: "readdir", Args: []Arg{Arg{Name: "dirp", Type: type_uint, Const: false}, Arg{Name: "entry", Type: &type_uint8, Const: false}, Arg{Name: "result", Type: type_uint, Const: false}}},
	&Signature{Id: 90, Name: "mmap", Args: []Arg{Arg{Name: "", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 91, Name: "munmap", Args: []Arg{Arg{Name: "addr", Type: type_uint32, Const: false}, Arg{Name: "len", Type: type_uint64, Const: false}}},
	&Signature{Id: 92, Name: "truncate", Args: []Arg{Arg{Name: "path", Type: type_stringc, Const: true}, Arg{Name: "length", Type: type_uint32, Const: false}}},
	&Signature{Id: 93, Name: "ftruncate", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "length", Type: type_uint32, Const: false}}},
	&Signature{Id: 94, Name: "fchmod", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "mode", Type: type_uint16, Const: false}}},
	&Signature{Id: 95, Name: "fchown", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "user", Type: type_uint32, Const: false}, Arg{Name: "group", Type: type_uint32, Const: false}}},
	&Signature{Id: 96, Name: "getpriority", Args: []Arg{Arg{Name: "which", Type: type_int, Const: false}, Arg{Name: "who", Type: type_int, Const: false}}},
	&Signature{Id: 97, Name: "setpriority", Args: []Arg{Arg{Name: "which", Type: type_int, Const: false}, Arg{Name: "who", Type: type_int, Const: false}, Arg{Name: "niceval", Type: type_int, Const: false}}},
	&unknownSignature, // 97
	&Signature{Id: 99, Name: "statfs", Args: []Arg{Arg{Name: "pathname", Type: type_stringc, Const: true}, Arg{Name: "buf", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 100, Name: "fstatfs", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "buf", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 101, Name: "ioperm", Args: []Arg{Arg{Name: "from", Type: type_uint32, Const: false}, Arg{Name: "num", Type: type_uint32, Const: false}, Arg{Name: "turn_on", Type: type_int, Const: false}}},
	&Signature{Id: 102, Name: "socketcall", Args: []Arg{Arg{Name: "call", Type: type_int, Const: false}, Arg{Name: "args", Type: &type_uint32, Const: false}}},
	&Signature{Id: 103, Name: "syslog", Args: []Arg{Arg{Name: "type", Type: type_int, Const: false}, Arg{Name: "buf", Type: type_stringc, Const: false}, Arg{Name: "len", Type: type_int, Const: false}}},
	&Signature{Id: 104, Name: "setitimer", Args: []Arg{Arg{Name: "which", Type: type_int, Const: false}, Arg{Name: "value", Type: &type_unknownstruct, Const: false}, Arg{Name: "ovalue", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 105, Name: "getitimer", Args: []Arg{Arg{Name: "which", Type: type_int, Const: false}, Arg{Name: "value", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 106, Name: "newstat", Args: []Arg{Arg{Name: "", Type: type_stringc, Const: false}, Arg{Name: "", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 107, Name: "newlstat", Args: []Arg{Arg{Name: "", Type: type_stringc, Const: false}, Arg{Name: "", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 108, Name: "newfstat", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "stat", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 109, Name: "uname", Args: []Arg{Arg{Name: "name", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 110, Name: "iopl", Args: []Arg{Arg{Name: "level", Type: type_uint32, Const: false}}},
	&Signature{Id: 111, Name: "vhangup", Args: []Arg{}},
	&Signature{Id: 112, Name: "idle", Args: []Arg{}},
	&Signature{Id: 113, Name: "vm86old", Args: []Arg{Arg{Name: "fn", Type: type_uint32, Const: false}, Arg{Name: "v86", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 114, Name: "wait4", Args: []Arg{Arg{Name: "upid", Type: type_int32, Const: false}, Arg{Name: "stat_addr", Type: &type_uint32, Const: false}, Arg{Name: "options", Type: type_int, Const: false}, Arg{Name: "ru", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 115, Name: "swapoff", Args: []Arg{Arg{Name: "specialfile", Type: type_stringc, Const: true}}},
	&Signature{Id: 116, Name: "sysinfo", Args: []Arg{Arg{Name: "info", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 117, Name: "ipc uint call", Args: []Arg{Arg{Name: "first", Type: type_int, Const: false}, Arg{Name: "second", Type: type_int, Const: false}, Arg{Name: "third", Type: type_int, Const: false}, Arg{Name: "ptr", Type: &type_uint8, Const: false}}},
	&Signature{Id: 118, Name: "fsync", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}}},
	&Signature{Id: 119, Name: "sigreturn", Args: []Arg{Arg{Name: "unused", Type: type_uint32, Const: false}}},
	&Signature{Id: 120, Name: "clone", Args: []Arg{Arg{Name: "clone_flags", Type: type_unknownstruct, Const: false}}},
	&Signature{Id: 121, Name: "setdomainname", Args: []Arg{Arg{Name: "name", Type: type_stringc, Const: false}, Arg{Name: "len", Type: type_int, Const: false}}},
	&Signature{Id: 122, Name: "newuname", Args: []Arg{Arg{Name: "", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 123, Name: "modify_ldt", Args: []Arg{Arg{Name: "func", Type: type_int, Const: false}, Arg{Name: "ptr", Type: &type_uint8, Const: false}, Arg{Name: "bytecount", Type: type_uint32, Const: false}}},
	&Signature{Id: 124, Name: "adjtimex", Args: []Arg{Arg{Name: "txc_p", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 125, Name: "mprotect", Args: []Arg{Arg{Name: "start", Type: type_uint32, Const: false}, Arg{Name: "len", Type: type_uint64, Const: false}, Arg{Name: "prot", Type: type_uint32, Const: false}}},
	&Signature{Id: 126, Name: "sigprocmask", Args: []Arg{Arg{Name: "how", Type: type_int, Const: false}, Arg{Name: "set", Type: &type_unknownstruct, Const: false}, Arg{Name: "oldset", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 127, Name: "create_module", Args: []Arg{Arg{Name: "name", Type: type_stringc, Const: true}, Arg{Name: "size", Type: type_uint64, Const: false}}},
	&Signature{Id: 128, Name: "init_module", Args: []Arg{Arg{Name: "umod", Type: type_stringc, Const: true}, Arg{Name: "len", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 129, Name: "delete_module", Args: []Arg{Arg{Name: "name_user", Type: type_stringc, Const: true}}},
	&Signature{Id: 130, Name: "get_kernel_syms", Args: []Arg{Arg{Name: "", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 131, Name: "quotactl", Args: []Arg{Arg{Name: "cmd", Type: type_int, Const: false}, Arg{Name: "special", Type: type_stringc, Const: true}, Arg{Name: "id", Type: type_int, Const: false}, Arg{Name: "addr", Type: type_uintptr, Const: false}}},
	&Signature{Id: 132, Name: "getpgid", Args: []Arg{Arg{Name: "pid", Type: type_int32, Const: false}}},
	&Signature{Id: 133, Name: "fchdir", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}}},
	&Signature{Id: 134, Name: "bdflush", Args: []Arg{Arg{Name: "func", Type: type_int, Const: false}, Arg{Name: "data", Type: type_int32, Const: false}}},
	&Signature{Id: 135, Name: "sysfs", Args: []Arg{Arg{Name: "option", Type: type_int, Const: false}, Arg{Name: "arg1", Type: type_uint32, Const: false}, Arg{Name: "arg2", Type: type_uint32, Const: false}}},
	&Signature{Id: 136, Name: "personality", Args: []Arg{Arg{Name: "personality", Type: type_uint32, Const: false}}},
	&unknownSignature, // 136
	&Signature{Id: 138, Name: "setfsuid", Args: []Arg{Arg{Name: "uid", Type: type_uint32, Const: false}}},
	&Signature{Id: 139, Name: "setfsgid", Args: []Arg{Arg{Name: "gid", Type: type_uint32, Const: false}}},
	&Signature{Id: 140, Name: "llseek", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "offset_high", Type: type_uint32, Const: false}, Arg{Name: "offset_low", Type: type_uint32, Const: false}, Arg{Name: "result", Type: &type_uint64, Const: false}, Arg{Name: "whence", Type: type_uint, Const: false}}},
	&Signature{Id: 141, Name: "getdents", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "dirent", Type: &type_uint8, Const: false}, Arg{Name: "count", Type: type_uint, Const: false}}},
	&Signature{Id: 142, Name: "select", Args: []Arg{Arg{Name: "n", Type: type_int, Const: false}, Arg{Name: "inp", Type: &type_unknownstruct, Const: false}, Arg{Name: "outp", Type: &type_unknownstruct, Const: false}, Arg{Name: "exp", Type: &type_unknownstruct, Const: false}, Arg{Name: "tvp", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 143, Name: "flock", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "cmd", Type: type_uint, Const: false}}},
	&Signature{Id: 144, Name: "msync", Args: []Arg{Arg{Name: "start", Type: type_uint32, Const: false}, Arg{Name: "len", Type: type_uint64, Const: false}, Arg{Name: "flags", Type: type_int, Const: false}}},
	&Signature{Id: 145, Name: "readv", Args: []Arg{Arg{Name: "fd", Type: type_uint32, Const: false}, Arg{Name: "vec", Type: &type_unknownstruct, Const: true}, Arg{Name: "vlen", Type: type_uint32, Const: false}}},
	&Signature{Id: 146, Name: "writev", Args: []Arg{Arg{Name: "fd", Type: type_uint32, Const: false}, Arg{Name: "vec", Type: &type_unknownstruct, Const: true}, Arg{Name: "vlen", Type: type_uint32, Const: false}}},
	&Signature{Id: 147, Name: "getsid", Args: []Arg{Arg{Name: "pid", Type: type_int32, Const: false}}},
	&Signature{Id: 148, Name: "fdatasync", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}}},
	&Signature{Id: 149, Name: "sysctl", Args: []Arg{Arg{Name: "", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 150, Name: "mlock", Args: []Arg{Arg{Name: "start", Type: type_uint32, Const: false}, Arg{Name: "len", Type: type_uint64, Const: false}}},
	&Signature{Id: 151, Name: "munlock", Args: []Arg{Arg{Name: "start", Type: type_uint32, Const: false}, Arg{Name: "len", Type: type_uint64, Const: false}}},
	&Signature{Id: 152, Name: "mlockall", Args: []Arg{Arg{Name: "flags", Type: type_int, Const: false}}},
	&Signature{Id: 153, Name: "munlockall", Args: []Arg{}},
	&Signature{Id: 154, Name: "sched_setparam", Args: []Arg{Arg{Name: "pid", Type: type_int32, Const: false}, Arg{Name: "param", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 155, Name: "sched_getparam", Args: []Arg{Arg{Name: "pid", Type: type_int32, Const: false}, Arg{Name: "param", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 156, Name: "sched_setscheduler", Args: []Arg{Arg{Name: "pid", Type: type_int32, Const: false}, Arg{Name: "policy", Type: type_int, Const: false}, Arg{Name: "param", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 157, Name: "sched_getscheduler", Args: []Arg{Arg{Name: "pid", Type: type_int32, Const: false}}},
	&Signature{Id: 158, Name: "sched_yield", Args: []Arg{}},
	&Signature{Id: 159, Name: "sched_get_priority_max", Args: []Arg{Arg{Name: "policy", Type: type_int, Const: false}}},
	&Signature{Id: 160, Name: "sched_get_priority_min", Args: []Arg{Arg{Name: "policy", Type: type_int, Const: false}}},
	&Signature{Id: 161, Name: "sched_rr_get_interval", Args: []Arg{Arg{Name: "pid", Type: type_int32, Const: false}, Arg{Name: "interval", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 162, Name: "nanosleep", Args: []Arg{Arg{Name: "rqtp", Type: &type_unknownstruct, Const: false}, Arg{Name: "rmtp", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 163, Name: "mremap", Args: []Arg{Arg{Name: "addr", Type: type_uint32, Const: false}, Arg{Name: "old_len", Type: type_uint32, Const: false}, Arg{Name: "new_len", Type: type_uint32, Const: false}, Arg{Name: "flags", Type: type_uint32, Const: false}}},
	&Signature{Id: 164, Name: "setresuid", Args: []Arg{Arg{Name: "ruid", Type: type_uint32, Const: false}, Arg{Name: "euid", Type: type_uint32, Const: false}, Arg{Name: "suid", Type: type_uint32, Const: false}}},
	&Signature{Id: 165, Name: "getresuid", Args: []Arg{Arg{Name: "ruid", Type: &type_uint32, Const: false}, Arg{Name: "euid", Type: &type_uint32, Const: false}, Arg{Name: "suid", Type: &type_uint32, Const: false}}},
	&Signature{Id: 166, Name: "vm86", Args: []Arg{Arg{Name: "", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 167, Name: "query_module", Args: []Arg{Arg{Name: "name", Type: type_stringc, Const: true}, Arg{Name: "which", Type: type_int, Const: false}, Arg{Name: "buf", Type: type_stringc, Const: false}, Arg{Name: "bufsize", Type: type_uint64, Const: false}, Arg{Name: "ret", Type: &type_uint64, Const: false}}},
	&Signature{Id: 168, Name: "poll", Args: []Arg{Arg{Name: "ufds", Type: &type_unknownstruct, Const: false}, Arg{Name: "nfds", Type: type_uint, Const: false}, Arg{Name: "timeout_msecs", Type: type_int32, Const: false}}},
	&Signature{Id: 169, Name: "nfsservctl", Args: []Arg{Arg{Name: "cmd", Type: type_int, Const: false}, Arg{Name: "argp", Type: &type_uint8, Const: false}, Arg{Name: "resp", Type: &type_uint8, Const: false}}},
	&Signature{Id: 170, Name: "setresgid", Args: []Arg{Arg{Name: "rgid", Type: type_uint32, Const: false}, Arg{Name: "egid", Type: type_uint32, Const: false}, Arg{Name: "sgid", Type: type_uint32, Const: false}}},
	&Signature{Id: 171, Name: "getresgid", Args: []Arg{Arg{Name: "rgid", Type: &type_uint32, Const: false}, Arg{Name: "egid", Type: &type_uint32, Const: false}, Arg{Name: "sgid", Type: &type_uint32, Const: false}}},
	&Signature{Id: 172, Name: "prctl", Args: []Arg{Arg{Name: "option", Type: type_int, Const: false}, Arg{Name: "arg2", Type: type_uint32, Const: false}, Arg{Name: "arg3", Type: type_uint32, Const: false}, Arg{Name: "arg4", Type: type_uint32, Const: false}, Arg{Name: "arg5", Type: type_uint32, Const: false}}},
	&Signature{Id: 173, Name: "rt_sigreturn", Args: []Arg{Arg{Name: "__unused", Type: type_uint32, Const: false}}},
	&Signature{Id: 174, Name: "rt_sigaction", Args: []Arg{Arg{Name: "sig", Type: type_int, Const: false}, Arg{Name: "act", Type: &type_unknownstruct, Const: true}, Arg{Name: "oact", Type: &type_unknownstruct, Const: false}, Arg{Name: "sigsetsize", Type: type_uint64, Const: false}}},
	&Signature{Id: 175, Name: "rt_sigprocmask", Args: []Arg{Arg{Name: "how", Type: type_int, Const: false}, Arg{Name: "nset", Type: &type_unknownstruct, Const: false}, Arg{Name: "oset", Type: &type_unknownstruct, Const: false}, Arg{Name: "sigsetsize", Type: type_uint64, Const: false}}},
	&Signature{Id: 176, Name: "rt_sigpending", Args: []Arg{Arg{Name: "set", Type: &type_unknownstruct, Const: false}, Arg{Name: "sigsetsize", Type: type_uint64, Const: false}}},
	&Signature{Id: 177, Name: "rt_sigtimedwait", Args: []Arg{Arg{Name: "uthese", Type: &type_unknownstruct, Const: true}, Arg{Name: "uinfo", Type: &type_unknownstruct, Const: false}, Arg{Name: "uts", Type: &type_unknownstruct, Const: true}, Arg{Name: "sigsetsize", Type: type_uint64, Const: false}}},
	&Signature{Id: 178, Name: "rt_sigqueueinfo", Args: []Arg{Arg{Name: "pid", Type: type_int, Const: false}, Arg{Name: "sig", Type: type_int, Const: false}, Arg{Name: "uinfo", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 179, Name: "rt_sigsuspend", Args: []Arg{Arg{Name: "unewset", Type: &type_unknownstruct, Const: false}, Arg{Name: "sigsetsize", Type: type_uint64, Const: false}}},
	&Signature{Id: 180, Name: "pread", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "buf", Type: type_stringc, Const: false}, Arg{Name: "count", Type: type_uint64, Const: false}, Arg{Name: "offset", Type: type_uint64, Const: false}}},
	&Signature{Id: 181, Name: "pwrite", Args: []Arg{Arg{Name: "fd", Type: type_uint, Const: false}, Arg{Name: "but", Type: type_stringc, Const: true}, Arg{Name: "count", Type: type_uint64, Const: false}, Arg{Name: "offset", Type: type_uint64, Const: false}}},
	&Signature{Id: 182, Name: "chown", Args: []Arg{Arg{Name: "filename", Type: type_stringc, Const: true}, Arg{Name: "user", Type: type_uint32, Const: false}, Arg{Name: "group", Type: type_uint32, Const: false}}},
	&Signature{Id: 183, Name: "getcwd", Args: []Arg{Arg{Name: "buf", Type: type_stringc, Const: false}, Arg{Name: "size", Type: type_uint32, Const: false}}},
	&Signature{Id: 184, Name: "capget", Args: []Arg{Arg{Name: "header", Type: type_unknownstruct, Const: false}, Arg{Name: "dataptr", Type: type_unknownstruct, Const: false}}},
	&Signature{Id: 185, Name: "capset", Args: []Arg{Arg{Name: "header", Type: type_unknownstruct, Const: false}, Arg{Name: "data", Type: type_unknownstruct, Const: true}}},
	&Signature{Id: 186, Name: "sigaltstack", Args: []Arg{Arg{Name: "uss", Type: &type_unknownstruct, Const: true}, Arg{Name: "uoss", Type: &type_unknownstruct, Const: false}}},
	&Signature{Id: 187, Name: "sendfile", Args: []Arg{Arg{Name: "out_fd", Type: type_int, Const: false}, Arg{Name: "in_fd", Type: type_int, Const: false}, Arg{Name: "offset", Type: &type_uint32, Const: false}, Arg{Name: "count", Type: type_uint64, Const: false}}},
	&unknownSignature, // 187
	&unknownSignature, // 188
	&Signature{Id: 190, Name: "vfork", Args: []Arg{Arg{Name: "regs", Type: &type_unknownstruct, Const: false}}},
}
