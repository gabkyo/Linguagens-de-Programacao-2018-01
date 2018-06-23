(in-ns 'lp4.main)

(defprotocol Lote
	(Id [this])
	(Owner [this])
	(Categoria [this])
	(Preco [this])
	(Area [this])
	(Info [this]))

(defrecord Terreno[id owner categoria solo pm2 base1 base2 altura]
	Lote
	(Id [this]
		(:id this))
	(Owner [this]
		(:owner this))
	(Categoria [this]
		(:categoria this))
	(Area [this]
		(case (:categoria this)
			"retang" (* (:base1 this) (:altura this))
			"triang" (/ (* (:base1 this) (:altura this)) 2)
			"trapez" (* (/ (+ (:base1 this) (:base2 this)) 2) (:altura this))
			0))
	(Preco [this]
		(case (:solo this)
			"A" (* (* (:pm2 this) (Area this)) 0.9)
			"R" (* (* (:pm2 this) (Area this)) 1.1)
			"G" (* (* (:pm2 this) (Area this)) 1.3)
			0))
	(Info [this]
		(:solo this))
)