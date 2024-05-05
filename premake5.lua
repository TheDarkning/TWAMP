include "dependencies/conandeps.premake5.lua"

workspace "TWAMP"
    configurations { "Debug", "Release" }
    architecture "x64"

    project "TWAMP"
        kind "ConsoleApp"
        language "C++"
        cppdialect "C++20"
        
        targetdir   "build/%{cfg.buildcfg}/bin"
        objdir      "build/%{cfg.buildcfg}/obj"

        location "./src"
        
        debugdir "app"

        linkoptions { conan_exelinkflags }

        files { "**.h", "**.cpp" }

        filter "configurations:Debug"
            defines { "DEBUG" }
            symbols "On"
        filter {}

        filter "configurations:Release"
            defines { "NDEBUG" }
            optimize "On"
        filter {}

        conan_setup()
        linkoptions { "/IGNORE:4099" }