package npcs

import (
	"fmt"
	"math"
)

func (location Location) String() string  {
	return fmt.Sprintf("(%f,%f,%f)",location.x, location.y, location.z)
}

/*
获取两个点的距离
 */
func (loc Location) EulideanDistance(target Location) float64  {

	return math.Sqrt((loc.x-target.x)*(loc.x-target.x)+(loc.y-target.y)*(loc.y-target.y)+(loc.z-target.z)*(loc.z-target.z))
}

func (npc NonPlayerCharacter) DistanceTo(character NonPlayerCharacter) float64 {

	return npc.Loc.EulideanDistance(character.Loc)
}

func (npc NonPlayerCharacter) String() string{

	return fmt.Sprintf("%s %s", npc.Name, npc.Loc)
}