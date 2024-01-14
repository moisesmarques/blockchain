// Copyright © 2015 Nik Unger
//
// This file is part of ringsig.
//
// Ringsig is free software: you can redistribute it and/or modify it under the
// terms of the GNU Lesser General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// Ringsig is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more
// details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with ringsig. If not, see <http://www.gnu.org/licenses/>.

//https://formulae.brew.sh/formula/pbc-sig#default
//https://formulae.brew.sh/formula/gmp#default

package main

import (
	"github.com/Nik-U/ringsig"
	"github.com/Nik-U/ringsig/shacham"
)


// A ring signature is a digital signature that is created by a member of a group which each have their own keys. 
// It is then not be possible to determine the person in the group who has created the signature.
//  A ring signature provides anonymity, unforgivably and collusion resistance.

func main() {
	// Generate new scheme parameters
	scheme, _ := shacham.New(2)

	// In the real world, these parameters would be saved using scheme.WriteTo
	// and then loaded by the clients using shacham.Load.

	// Two clients generate key pairs
	cpuSubKey := scheme.KeyGen()
	gpuSubKey := scheme.KeyGen()

	// We will sign over the ring of the two users. In general, higher level
	// communication protocols will somehow specify the ring used to sign the
	// message (either explicitly or implicitly).
	ring := []ringsig.PublicKey{cpuSubKey.Public, gpuSubKey.Public}

	sig, _ := scheme.Sign( "which data ?", ring, cpuSubKey) // to sign transactions

	if scheme.Verify("which data ?", sig, ring) { } // who should verify this data? the shell?
}
