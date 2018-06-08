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
			  (reset! categoria c)))
)