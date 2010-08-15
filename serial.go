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
	"os"
	"termios"
)

type Serial struct {
	*os.File
	ot termios.Termios
}

const (
	B9600_8E2 = termios.B9600 | termios.CS8 | termios.PARENB | termios.CSTOPB
)

func Open(dev string, flag int, perm uint32, cflags uint) (*Serial, os.Error) {
	f, err := os.Open(dev, flag, perm)
	if err != nil {
		return nil, err
	}

	fd := f.Fd()
	s := Serial{File: f}

	if err = s.ot.Get(fd); err != nil {
		return nil, err
	}

	var t termios.Termios

	t.CFlag = termios.CREAD | termios.HUPCL | termios.CLOCAL | cflags

	// TODO: figure out why this works and if
	// other entries in CC should be set
	t.CC[termios.VMIN] = 1

	if err = t.Set(fd); err != nil {
		return nil, err
	}

	return &s, nil
}

func (s *Serial) Close() os.Error {
	s.ot.Set(s.Fd())

	return s.File.Close()
}
