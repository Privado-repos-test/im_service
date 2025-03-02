/**
 * Copyright (c) 2014-2015, GoBelieve
 * All rights reserved.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 59 Temple Place, Suite 330, Boston, MA  02111-1307  USA
 */

package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type RedisConfig struct {
	Address  string `toml:"address"`
	Password string `toml:"password"`
	Db       int    `toml:"db"`
}

type LogConfig struct {
	Filename string `toml:"filename"`
	Level    string `toml:"level"`
	Backup   int    `toml:"backup"` //log files
	Age      int    `toml:"age"`    //days
	Caller   bool   `toml:"caller"`
}

type Config struct {
	Listen string `toml:"listen"`

	PushDisabled      bool   `toml:"push_disabled"`
	HttpListenAddress string `toml:"http_listen_address"`

	Redis RedisConfig `toml:"redis"`
	Log   LogConfig   `toml:"log"`
}

func read_route_cfg(cfg_path string) *Config {
	var conf Config
	if _, err := toml.DecodeFile(cfg_path, &conf); err != nil {
		// handle error
		log.Fatal("Decode cfg file fail:", err)
	}
	return &conf
}
