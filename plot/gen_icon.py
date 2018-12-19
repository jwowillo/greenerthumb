"""
gen_icon.py is a utility script that generates the icon into the correct c++
file.

It should be run whenever the icon is updated.
"""

FILE = """
#pragma once

#include <cstddef>
#include <vector>

// Generated file containing 'icon.png'.

namespace render {{

const std::vector<uint8_t> ICON{{{}}};

}} //  namespace render"""

bs = []

with open('icon.png', 'rb') as f:
    byte = f.read(1)
    while byte != b'':
        bs.append(hex(ord(byte)))
        byte = f.read(1)

print(FILE.format(','.join(bs)))
