(in-ns 'lp4.main)

(defn sortImoveis [a b]
	(compare [(Preco b) (Id a)] [(Preco a) (Id b)])) ;;retorna qual tem o maior preco ou menor id se preco eh igual

(defn testar [a]
	(pop a))