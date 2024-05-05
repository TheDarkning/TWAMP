#!lua

include "conanutils.premake5.lua"

t_conandeps = {}
t_conandeps["debug_x86_64"] = {}
t_conandeps["debug_x86_64"]["inja"] = {}
t_conandeps["debug_x86_64"]["inja"]["includedirs"] = {"C:/Users/erinzha/.conan2/p/inja62ccb399ee74d/p/include"}
t_conandeps["debug_x86_64"]["inja"]["libdirs"] = {}
t_conandeps["debug_x86_64"]["inja"]["bindirs"] = {}
t_conandeps["debug_x86_64"]["inja"]["libs"] = {}
t_conandeps["debug_x86_64"]["inja"]["system_libs"] = {}
t_conandeps["debug_x86_64"]["inja"]["defines"] = {}
t_conandeps["debug_x86_64"]["inja"]["cxxflags"] = {}
t_conandeps["debug_x86_64"]["inja"]["cflags"] = {}
t_conandeps["debug_x86_64"]["inja"]["sharedlinkflags"] = {}
t_conandeps["debug_x86_64"]["inja"]["exelinkflags"] = {}
t_conandeps["debug_x86_64"]["inja"]["frameworks"] = {}

if conandeps == nil then conandeps = {} end
conan_premake_tmerge(conandeps, t_conandeps)
