(in-ns 'lp4.main)

(defn sortImoveis [a b]
	(compare [(Preco a) (Id a)] [(Preco b) (Id b)])) 
	;;preco crescente id decrescente


(defn remover [lista id]
	(remove (fn [x] (= (Id x) id)) lista))

(defn criarCasas [lista plimite alimite] ;;retorna lista com casas com preco < plim e area >alimite
	(filter (fn [x] (and (and (= (Categoria x) "casa") (< (Preco x) plimite)) (> (Area x) alimite))) lista))

(defn sortCasas [a b]
	(compare [(Info b) (Id b)] [(Info a) (Id a)])) 
	;;area decrescente id crescente

(defn criarArgilosos [lista]
	(filter (fn [x] (= (Info x) "G")) lista))

(defn sortArgiloso [a b]
	(compare [(Area b) (Id b)] [(Area a) (Id a)])) 