#!lua

include "conanutils.premake5.lua"

t_conandeps = {}
t_conandeps["release_x86_64"] = {}
t_conandeps["release_x86_64"]["xxhash"] = {}
t_conandeps["release_x86_64"]["xxhash"]["includedirs"] = {"C:/Users/erinzha/.conan2/p/xxhascaafe0666c80b/p/include"}
t_conandeps["release_x86_64"]["xxhash"]["libdirs"] = {"C:/Users/erinzha/.conan2/p/xxhascaafe0666c80b/p/lib"}
t_conandeps["release_x86_64"]["xxhash"]["bindirs"] = {"C:/Users/erinzha/.conan2/p/xxhascaafe0666c80b/p/bin"}
t_conandeps["release_x86_64"]["xxhash"]["libs"] = {"xxhash"}
t_conandeps["release_x86_64"]["xxhash"]["system_libs"] = {}
t_conandeps["release_x86_64"]["xxhash"]["defines"] = {}
t_conandeps["release_x86_64"]["xxhash"]["cxxflags"] = {}
t_conandeps["release_x86_64"]["xxhash"]["cflags"] = {}
t_conandeps["release_x86_64"]["xxhash"]["sharedlinkflags"] = {}
t_conandeps["release_x86_64"]["xxhash"]["exelinkflags"] = {}
t_conandeps["release_x86_64"]["xxhash"]["frameworks"] = {}

if conandeps == nil then conandeps = {} end
conan_premake_tmerge(conandeps, t_conandeps)
