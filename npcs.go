package npcs

import (
	"fmt"
	"math"
)

func (location Location) String() string  {
	return fmt.Sprintf("(%f,%f,%f)",location.X, location.Y, location.Z)
}

/*
获取两个点的距离
 */
func (loc Location) EulideanDistance(target Location) float64  {

	return math.Sqrt((loc.X-target.X)*(loc.X-target.X)+(loc.Y-target.Y)*(loc.Y-target.Y)+(loc.Z-target.Z)*(loc.Z-target.Z))
}

func (npc NonPlayerCharacter) DistanceTo(character NonPlayerCharacter) float64 {

	return npc.Loc.EulideanDistance(character.Loc)
}

func (npc NonPlayerCharacter) String() string{

	return fmt.Sprintf("%s %s", npc.Name, npc.Loc)
}