(ns classes)

(defprotocol Lote
	"Base"
	(Preco [])
	(Info [])
	(Area [])
	(Id [])
	(Categoria [])
	(Owner []))

(defrecord Terreno [] Lote
	(let [
    	id (atom nil)
		owner (atom nil)
		categoria (atom nil)
		solo (atom nil)
		precom2(atom nil)
		base1 (atom nil)
		base2 (atom nil)
		altura (atom nil)
	])
	set (fn [i o c s p b1 b2 a]
          (do (reset! id i)
			  (reset! owner o)
			  (reset! categoria c)
			  (reset! solo s)
			  (reset! precom2 p)
			  (reset! base1 b1)
			  (reset! base2 b2)
			  (reset! altura a)
			  ))

	(Area []
		(if (= :categoria "triang")
			(/ (* :base1 :altura) 2)
			(if (= :categoria "retang")
				(* :base1 :altura) 
				(if (= :categoria "trapez")
					(/ (* (+ base1 base2) altura) 2)
					0))))

	(Preco [] 
		

	)
)