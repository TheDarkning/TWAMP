#!lua

include "conanutils.premake5.lua"

t_conandeps = {}
t_conandeps["debug_x86_64"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"]["includedirs"] = {"C:/Users/erinzha/.conan2/p/cpp-hb43040248388d/p/include",
"C:/Users/erinzha/.conan2/p/cpp-hb43040248388d/p/include/httplib"}
t_conandeps["debug_x86_64"]["cpp-httplib"]["libdirs"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"]["bindirs"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"]["libs"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"]["system_libs"] = {"crypt32", "cryptui", "ws2_32"}
t_conandeps["debug_x86_64"]["cpp-httplib"]["defines"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"]["cxxflags"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"]["cflags"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"]["sharedlinkflags"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"]["exelinkflags"] = {}
t_conandeps["debug_x86_64"]["cpp-httplib"]["frameworks"] = {}

if conandeps == nil then conandeps = {} end
conan_premake_tmerge(conandeps, t_conandeps)
