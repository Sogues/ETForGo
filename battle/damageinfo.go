package battle

type (
	Damage struct {
		bullet    int
		explosion int
		mental    int
	}

	SDamageInfo struct {
		attacker GameObject

		defender GameObject

		tags []DamageInfoTag

		damage Damage

		criticalRate float64

		hitRate float64

		//伤害的角度，作为伤害打向角色的入射角度，比如子弹，就是它当前的飞行角度
		//潘森抵抗正面的伤害
		degree float64

		addBuffs map[*AddBuffInfo]struct{}
	}
	DamageInfo *SDamageInfo
)