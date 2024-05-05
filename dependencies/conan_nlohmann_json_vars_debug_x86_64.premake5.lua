#!lua

include "conanutils.premake5.lua"

t_conandeps = {}
t_conandeps["debug_x86_64"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["includedirs"] = {"C:/Users/erinzha/.conan2/p/nlohm0567ffc90cfc1/p/include"}
t_conandeps["debug_x86_64"]["nlohmann_json"]["libdirs"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["bindirs"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["libs"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["system_libs"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["defines"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["cxxflags"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["cflags"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["sharedlinkflags"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["exelinkflags"] = {}
t_conandeps["debug_x86_64"]["nlohmann_json"]["frameworks"] = {}

if conandeps == nil then conandeps = {} end
conan_premake_tmerge(conandeps, t_conandeps)
