#!lua

include "conanutils.premake5.lua"

t_conandeps = {}
t_conandeps["release_x86_64"] = {}
t_conandeps["release_x86_64"]["jsoncpp"] = {}
t_conandeps["release_x86_64"]["jsoncpp"]["includedirs"] = {"C:/Users/erinzha/.conan2/p/jsonceac8e8b48ba79/p/include"}
t_conandeps["release_x86_64"]["jsoncpp"]["libdirs"] = {"C:/Users/erinzha/.conan2/p/jsonceac8e8b48ba79/p/lib"}
t_conandeps["release_x86_64"]["jsoncpp"]["bindirs"] = {"C:/Users/erinzha/.conan2/p/jsonceac8e8b48ba79/p/bin"}
t_conandeps["release_x86_64"]["jsoncpp"]["libs"] = {"jsoncpp"}
t_conandeps["release_x86_64"]["jsoncpp"]["system_libs"] = {}
t_conandeps["release_x86_64"]["jsoncpp"]["defines"] = {}
t_conandeps["release_x86_64"]["jsoncpp"]["cxxflags"] = {}
t_conandeps["release_x86_64"]["jsoncpp"]["cflags"] = {}
t_conandeps["release_x86_64"]["jsoncpp"]["sharedlinkflags"] = {}
t_conandeps["release_x86_64"]["jsoncpp"]["exelinkflags"] = {}
t_conandeps["release_x86_64"]["jsoncpp"]["frameworks"] = {}

if conandeps == nil then conandeps = {} end
conan_premake_tmerge(conandeps, t_conandeps)
