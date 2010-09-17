/*
 * This file is part of go-serial.
 *
 * go-serial is free software: you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation, either version 3 of
 * the License, or(at your option) any later version.
 *
 * go-serial is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with go-serial.  If not, see <http://www.gnu.org/licenses/>.
 */
package serial

import (
	"termios"
)

// speed constants
const (
	B9600 = termios.B9600
)

// mode constants
const (
	MODE_8E1 = termios.CS8 | termios.PARENB | termios.CSTOPB
)
