#!lua

include "conanutils.premake5.lua"

t_conandeps = {}
t_conandeps["debug_x86_64"] = {}
t_conandeps["debug_x86_64"]["xxhash"] = {}
t_conandeps["debug_x86_64"]["xxhash"]["includedirs"] = {"C:/Users/erinzha/.conan2/p/b/xxhas627d2a9d181ff/p/include"}
t_conandeps["debug_x86_64"]["xxhash"]["libdirs"] = {"C:/Users/erinzha/.conan2/p/b/xxhas627d2a9d181ff/p/lib"}
t_conandeps["debug_x86_64"]["xxhash"]["bindirs"] = {"C:/Users/erinzha/.conan2/p/b/xxhas627d2a9d181ff/p/bin"}
t_conandeps["debug_x86_64"]["xxhash"]["libs"] = {"xxhash"}
t_conandeps["debug_x86_64"]["xxhash"]["system_libs"] = {}
t_conandeps["debug_x86_64"]["xxhash"]["defines"] = {}
t_conandeps["debug_x86_64"]["xxhash"]["cxxflags"] = {}
t_conandeps["debug_x86_64"]["xxhash"]["cflags"] = {}
t_conandeps["debug_x86_64"]["xxhash"]["sharedlinkflags"] = {}
t_conandeps["debug_x86_64"]["xxhash"]["exelinkflags"] = {}
t_conandeps["debug_x86_64"]["xxhash"]["frameworks"] = {}

if conandeps == nil then conandeps = {} end
conan_premake_tmerge(conandeps, t_conandeps)
