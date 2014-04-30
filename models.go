/*
   Copyright 2013-2014 gtalent2@gmail.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/
package models

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"strings"
)


type CreatureType struct {
	Name map[string]string
	Special bool
	StrongAgainst []string
	WeakAgainst []string
}

func (me *CreatureType) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *CreatureType) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *CreatureType) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *CreatureType) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *CreatureType) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *CreatureType) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *CreatureType) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *CreatureType) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *CreatureType) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type StatusEffect struct {
	AttackerEffect int
	EnemyEffect int
}

func (me *StatusEffect) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *StatusEffect) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *StatusEffect) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *StatusEffect) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *StatusEffect) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *StatusEffect) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *StatusEffect) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *StatusEffect) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *StatusEffect) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Fraction struct {
	Current int
	Available int
}

func (me *Fraction) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Fraction) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Fraction) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Fraction) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Fraction) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Fraction) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Fraction) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Fraction) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Fraction) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type ModelFile struct {
	Type string
}

func (me *ModelFile) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *ModelFile) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *ModelFile) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *ModelFile) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *ModelFile) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *ModelFile) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *ModelFile) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *ModelFile) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *ModelFile) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Point struct {
	X int
	Y int
}

func (me *Point) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Point) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Point) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Point) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Point) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Point) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Point) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Point) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Point) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Size struct {
	Width int
	Height int
}

func (me *Size) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Size) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Size) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Size) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Size) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Size) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Size) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Size) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Size) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Bounds struct {
	X int
	Y int
	Width int
	Height int
}

func (me *Bounds) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Bounds) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Bounds) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Bounds) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Bounds) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Bounds) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Bounds) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Bounds) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Bounds) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type SaveVariables struct {
	Vars map[string]interface{}
}

func (me *SaveVariables) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *SaveVariables) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *SaveVariables) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *SaveVariables) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *SaveVariables) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *SaveVariables) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *SaveVariables) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *SaveVariables) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *SaveVariables) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type SpriteSheetImage struct {
	SrcBounds Bounds
}

func (me *SpriteSheetImage) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *SpriteSheetImage) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *SpriteSheetImage) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *SpriteSheetImage) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *SpriteSheetImage) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *SpriteSheetImage) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *SpriteSheetImage) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *SpriteSheetImage) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *SpriteSheetImage) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type SpriteSheet struct {
	TilesWide int
	TilesHigh int
	TileWidth int
	TileHeight int
	SrcFile string
	Images map[int]SpriteSheetImage
	ImageIdIterator int
	RecycledImageIds []int
}

func (me *SpriteSheet) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *SpriteSheet) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *SpriteSheet) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *SpriteSheet) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *SpriteSheet) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *SpriteSheet) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *SpriteSheet) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *SpriteSheet) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *SpriteSheet) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type CreatureMove struct {
	Name map[string]string
	Type string
	Power int
	RequiresRegarge bool
	Script string
	Burn StatusEffect
	Freeze StatusEffect
	Paralyze StatusEffect
	Poison StatusEffect
	Sleep StatusEffect
}

func (me *CreatureMove) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *CreatureMove) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *CreatureMove) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *CreatureMove) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *CreatureMove) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *CreatureMove) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *CreatureMove) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *CreatureMove) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *CreatureMove) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type CreatureMoveInstance struct {
	CreatureMove int
	PP Fraction
}

func (me *CreatureMoveInstance) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *CreatureMoveInstance) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *CreatureMoveInstance) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *CreatureMoveInstance) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *CreatureMoveInstance) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *CreatureMoveInstance) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *CreatureMoveInstance) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *CreatureMoveInstance) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *CreatureMoveInstance) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Creature struct {
	ID int
	Name map[string]string
	CreatureClass string
	Male bool
	Level int
	Health Fraction
	Attack int
	SpecAttack int
	Defense int
	SpecDefense int
	Burned bool
	Frozen bool
	Poisoned bool
	Asleep bool
	Moves []CreatureMoveInstance
}

func (me *Creature) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Creature) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Creature) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Creature) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Creature) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Creature) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Creature) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Creature) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Creature) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type User struct {
	PersonID int
	World string
	Zone string
	XAddress int
	YAddress int
	Layer int
}

func (me *User) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *User) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *User) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *User) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *User) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *User) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *User) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *User) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *User) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type ZoneInstance struct {
	AccessorID string
	ZonePath string
	Location Point
}

func (me *ZoneInstance) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *ZoneInstance) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *ZoneInstance) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *ZoneInstance) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *ZoneInstance) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *ZoneInstance) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *ZoneInstance) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *ZoneInstance) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *ZoneInstance) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type ZoneHeader struct {
	Path string
	Size Size
}

func (me *ZoneHeader) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *ZoneHeader) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *ZoneHeader) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *ZoneHeader) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *ZoneHeader) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *ZoneHeader) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *ZoneHeader) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *ZoneHeader) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *ZoneHeader) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type SaveFile struct {
	Vars map[string]interface{}
	User User
}

func (me *SaveFile) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *SaveFile) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *SaveFile) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *SaveFile) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *SaveFile) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *SaveFile) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *SaveFile) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *SaveFile) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *SaveFile) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Image struct {
	SpriteSheet string
	ImgId int
	DefaultSize Size
}

func (me *Image) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Image) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Image) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Image) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Image) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Image) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Image) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Image) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Image) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type World struct {
	Zones []ZoneInstance
}

func (me *World) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *World) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *World) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *World) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *World) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *World) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *World) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *World) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *World) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type AnimationSlide struct {
	Interval int
	Image Image
}

func (me *AnimationSlide) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *AnimationSlide) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *AnimationSlide) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *AnimationSlide) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *AnimationSlide) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *AnimationSlide) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *AnimationSlide) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *AnimationSlide) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *AnimationSlide) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Animation struct {
	Images []AnimationSlide
}

func (me *Animation) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Animation) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Animation) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Animation) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Animation) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Animation) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Animation) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Animation) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Animation) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type AnimLayer struct {
	Point Point
	Animation Animation
}

func (me *AnimLayer) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *AnimLayer) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *AnimLayer) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *AnimLayer) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *AnimLayer) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *AnimLayer) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *AnimLayer) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *AnimLayer) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *AnimLayer) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type CreatureClass struct {
	Name map[string]string
	Successor string
	Predecessor string
	Types []string
	CanLearn []string
	LearnsAtLevel map[int]string
	FrontView Animation
	BackView Animation
}

func (me *CreatureClass) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *CreatureClass) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *CreatureClass) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *CreatureClass) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *CreatureClass) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *CreatureClass) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *CreatureClass) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *CreatureClass) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *CreatureClass) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type PersonClass struct {
	ID int
	Name map[string]string
	Creatures []int
	Overhead []Animation
	FrontView Animation
	BackView Animation
}

func (me *PersonClass) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *PersonClass) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *PersonClass) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *PersonClass) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *PersonClass) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *PersonClass) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *PersonClass) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *PersonClass) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *PersonClass) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type TileClass struct {
	TerrainFlags int
	Import string
	LowerAnims []AnimLayer
	UpperAnims []AnimLayer
}

func (me *TileClass) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *TileClass) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *TileClass) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *TileClass) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *TileClass) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *TileClass) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *TileClass) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *TileClass) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *TileClass) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Person struct {
	PersonClass PersonClass
}

func (me *Person) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Person) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Person) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Person) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Person) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Person) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Person) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Person) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Person) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Sprite struct {
	AnimLayers [][]AnimLayer
	SpriteType int
	PersonID int
	Speed int
	Name string
	Path string
	ScriptPath string
}

func (me *Sprite) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Sprite) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Sprite) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Sprite) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Sprite) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Sprite) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Sprite) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Sprite) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Sprite) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Tile struct {
	TileClass TileClass
	Occupant Sprite
}

func (me *Tile) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Tile) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Tile) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Tile) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Tile) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Tile) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Tile) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Tile) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Tile) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}

type Zone struct {
	Tiles [][][]Tile
	InitScripts []string
	Location Point
}

func (me *Zone) FromJSON(text []byte) error {
	return json.Unmarshal(text, me)
}

func (me *Zone) ReadJSONFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return json.Unmarshal(file, me)
}

func (me *Zone) ToJSON() []byte {
	out, _ := json.Marshal(me)
	return out
}

func (me *Zone) WriteJSONFile(path string) error {
	out, _ := json.Marshal(me)
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Zone) ToGob() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	enc.Encode(me)
	return buf.Bytes()
}

func (me *Zone) WriteGobFile(path string) error {
	out := me.ToGob()
	return ioutil.WriteFile(path, out, 0644)
}

func (me *Zone) FromGob(data []byte) error {
	var buf bytes.Buffer
	_, err := buf.Write(data)
	if err != nil {
		return err
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(me)
	if err != nil {
		return err
	}
	return err
}

func (me *Zone) ReadGobFile(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		 return err
	}
	return me.FromGob(file)
}

func (me *Zone) ReadObjFile(path string) error {
	switch path[strings.LastIndex(path, ".")+1:] {
	case "json":
		return me.ReadJSONFile(path)
	case "gob":
		return me.ReadGobFile(path)
	default:
		err := me.ReadGobFile(path+".gob")
		if err == nil {
			return nil
		}
		err = me.ReadJSONFile(path+".json")
		if err == nil {
			return nil
		}
		return err
	}
	return nil
}
