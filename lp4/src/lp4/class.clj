(ns lp4.class)

(defprotocol Lote
	(Id [this])
	(Categoria [this])
	(Preco [this])
	(Area [this])
	(Info [this]))

(defrecord Terreno[id owner categoria solo pm2 base1 base2 altura]
	Lote
	(Id [this]
		(:id this))
	(Categoria [this]
		(:categoria this))
	(Area [this]
		(if (= categoria "retang")
			(* base1 altura)
			if (= categoria "triang")
				(/ (* base1 altura) 2)
				if (= categoria "trapez")
					(/ (* (+ base1 base2) altura) 2)
					0))
	(Preco [this]
		(if (= solo "A")))
)