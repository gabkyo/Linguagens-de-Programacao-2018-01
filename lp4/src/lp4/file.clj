(in-ns 'lp4.main)
(load "class")
(load "file")
(require '[clojure.string :as str])
(require '[clojure.java.io :as io])

(defn parse-int [s]
   (Integer. (re-find  #"\d+" s )))

(defn catalogo []
	(let [lista (seq nil)]
		(if (.exists (io/file "catalogo.txt")) 
		(do
			(def linhas (str/split (slurp "catalogo.txt") #"\n"))
			(while (not (empty? linhas)) 
				(do
				(def categoria (first linhas))
				(def linhas (rest linhas))
				(def id (Integer/parseInt (first linhas)))
				(def linhas (rest linhas))
				(def owner (first linhas))
				(def linhas (rest linhas))
				(case categoria
					"casa" (do
							()
						)
					"apto" (do

						) 
					"trapez" (do

						)
					"triang" (do

						)
					"retang" (do

						)
					(println "Categoria:" categoria))
				(def linhas (rest linhas))
				)
			)
		)
		(println "Catalogo nao encontrado."))
		lista
	)
)