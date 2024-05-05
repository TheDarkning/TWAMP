#!lua

include "conanutils.premake5.lua"

t_conandeps = {}
t_conandeps["release_x86_64"] = {}
t_conandeps["release_x86_64"]["openssl"] = {}
t_conandeps["release_x86_64"]["openssl"]["includedirs"] = {"C:/Users/erinzha/.conan2/p/opens14ab97bbf4437/p/include"}
t_conandeps["release_x86_64"]["openssl"]["libdirs"] = {"C:/Users/erinzha/.conan2/p/opens14ab97bbf4437/p/lib"}
t_conandeps["release_x86_64"]["openssl"]["bindirs"] = {"C:/Users/erinzha/.conan2/p/opens14ab97bbf4437/p/bin"}
t_conandeps["release_x86_64"]["openssl"]["libs"] = {"libssl", "libcrypto"}
t_conandeps["release_x86_64"]["openssl"]["system_libs"] = {"crypt32", "ws2_32", "advapi32", "user32", "bcrypt"}
t_conandeps["release_x86_64"]["openssl"]["defines"] = {}
t_conandeps["release_x86_64"]["openssl"]["cxxflags"] = {}
t_conandeps["release_x86_64"]["openssl"]["cflags"] = {}
t_conandeps["release_x86_64"]["openssl"]["sharedlinkflags"] = {}
t_conandeps["release_x86_64"]["openssl"]["exelinkflags"] = {}
t_conandeps["release_x86_64"]["openssl"]["frameworks"] = {}

if conandeps == nil then conandeps = {} end
conan_premake_tmerge(conandeps, t_conandeps)
