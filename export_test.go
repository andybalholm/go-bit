// Copyright 2012 Stefan Nilsson
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bit

func (S *Set) Sdata() []uint64             { return S.data }
func (S *Set) Smin() int                   { return S.min }
func FindMinFrom(n int, data []uint64) int { return findMinFrom(n, data) }
func NextPow2(n int) (p int)               { return nextPow2(n) }
