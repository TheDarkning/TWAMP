from conan import ConanFile
from conan.tools.build import check_min_cppstd
from conan.tools.files import copy, get
from conan.tools.layout import basic_layout
import os

class ExampleRecipe(ConanFile):
    settings = "os", "compiler", "build_type", "arch"
    generators = "PremakeDeps"

    options = {
        "with_openssl": [True, False],
        "with_zlib": [True, False],
        "with_brotli": [True, False],
    }

    default_options = {
        "with_openssl": True,
        "with_zlib": False,
        "with_brotli": False,
    }


    def requirements(self):
        self.requires("cpp-httplib/0.15.3")
        self.requires("openssl/[>=1.1 <4]")
        self.requires("jsoncpp/1.9.5")
        self.requires("inja/3.4.0")
        self.requires("xxhash/0.8.2")