#!lua

include "conanutils.premake5.lua"

t_conandeps = {}
t_conandeps["debug_x86_64"] = {}
t_conandeps["debug_x86_64"]["zlib"] = {}
t_conandeps["debug_x86_64"]["zlib"]["includedirs"] = {"C:/Users/erinzha/.conan2/p/b/zlib19bc09bb67b58/p/include"}
t_conandeps["debug_x86_64"]["zlib"]["libdirs"] = {"C:/Users/erinzha/.conan2/p/b/zlib19bc09bb67b58/p/lib"}
t_conandeps["debug_x86_64"]["zlib"]["bindirs"] = {"C:/Users/erinzha/.conan2/p/b/zlib19bc09bb67b58/p/bin"}
t_conandeps["debug_x86_64"]["zlib"]["libs"] = {"zlib"}
t_conandeps["debug_x86_64"]["zlib"]["system_libs"] = {}
t_conandeps["debug_x86_64"]["zlib"]["defines"] = {}
t_conandeps["debug_x86_64"]["zlib"]["cxxflags"] = {}
t_conandeps["debug_x86_64"]["zlib"]["cflags"] = {}
t_conandeps["debug_x86_64"]["zlib"]["sharedlinkflags"] = {}
t_conandeps["debug_x86_64"]["zlib"]["exelinkflags"] = {}
t_conandeps["debug_x86_64"]["zlib"]["frameworks"] = {}

if conandeps == nil then conandeps = {} end
conan_premake_tmerge(conandeps, t_conandeps)
