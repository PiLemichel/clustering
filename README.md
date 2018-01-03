# clustering
## Reflexion

## Algorithme

### Détails du Fonctionnement

- **Step 1** L’algorithme crée tout d’abords le maximum de cluster pour notre jeu de donnée. (10 dans notre cas)
- **Step 2** Puis calcule il calcule la distance de Hamming moyenne entre tous ces clusters.
- **Step 3** Compare les distances afin d’extraire la distance minimum entre deux clusters. (si égalité de des distances, il choisit la dernière        distance minimum rencontré) 
- **Step 4** Fusionne les clusters conserver.
- **Step 5** Recommencer depuis l’étape 2 jusqu’à ce que le nombre de cluster généré soit égal au nombre de cluster demandé 

### Exemple (procédure pour 6 clusters) 
 
Cluster de base à 10 elements (Step 1): 
```
[[{clair 2 2 fine}][{clair 1 2 fine}][{foncée 2 2 fine}] [{clair 2 1 fine}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse}] [{foncée 2 2 épaisse}] [{foncée 1 1 fine}] [{foncée 1 1 épaisse}] [{foncée 2 1 épaisse}]]
```
**{clair 2 2 fine}** et **{clair 1 2 fine}** pour une distance de 1 (Step 2 et Step 3), fusion des deux clusters (Step 4): 
```
[[{clair 2 2 fine} {clair 1 2 fine}]  [{foncée 2 2 fine}] [{clair 2 1 fine}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse}] [{foncée 2 2 épaisse}] [{foncée 1 1 fine}] [{foncée 1 1 épaisse}] [{foncée 2 1 épaisse}]]
```
**{foncée 2 2 fine}** et **{foncée 2 2 épaisse}** pour une distance de 1 (Step 2 et Step 3), fusion des deux clusters (Step 4): 
```
[[{clair 2 2 fine} {clair 1 2 fine}]  [{foncée 2 2 fine} {foncée 2 2 épaisse}] [{clair 2 1 fine}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse}] [{foncée 1 1 fine}] [{foncée 1 1 épaisse}] [{foncée 2 1 épaisse}]]
```
**{clair 1 1 épaisse}** et **{foncée 1 1 épaisse}** pour une distance de 1 (Step 2 et Step 3), fusion des deux clusters (Step 4):  
```
[[{clair 2 2 fine} {clair 1 2 fine}]  [{foncée 2 2 fine} {foncée 2 2 épaisse}] [{clair 2 1 fine}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse} {foncée 1 1 épaisse}]  [{foncée 1 1 fine}] [{foncée 2 1 épaisse}]]
```
**{clair 2 2 fine} {clair 1 2 fine}** et **{clair 2 1 fine}** 3 pour une distance de  1.5 (Step 2 et Step 3), fusion des deux clusters (Step 4): 
```
 [[{clair 2 2 fine} {clair 1 2 fine} {clair 2 1 fine}] [{foncée 2 2 fine} {foncée 2 2 épaisse}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse} {foncée 1 1 épaisse}]  [{foncée 1 1 fine}] [{foncée 2 1 épaisse}]]
```
## Commencer
### Prérequis
Installation de [Golang](https://golang.org/doc/install) 
### Mise en place
```
Installation de Golang sur votre machine: https://golang.org/doc/install
mkdir $HOME/go/src/github.com/PiLemichel/
cd $HOME/go/src/github.com/PiLemichel/
git clone git@github.com:PiLemichel/clustering.git
cd clustering
```
### Lancement
```
go build 
./cluster
```

## Retours

### Les axe d’amélioration 

-    L’algorithme ne fonction que pour un jeu de données, des 10 éléments composés eux même de 4 éléments. Un refacto’ du code afin de le rendre « adaptable » à N élément composé de N éléments est donc un axe d’amélioration majeur de cet algorithme.

-    Du fait du choix arbitraire de la distance minimal, l’algorithme ne fournit qu’une seule possibilité de cluster ce qui pourrait être remédié par la mise en place d’un choix « random » de la distance minimum. Nous obtiendrons ainsi des clusters différents à chaque lancement.

-    Sans les commentaire le code reste difficile à lire, un découpage plus important du code facilitera sans doute à la fois la lecture du code et serai par la même occasion une possibilité de l’optimiser ce qui n’ai pas le cas aujourd’hui.

### Colaborations
 Echange sur le fonctionnement que dois adopter l’algorithme : « Axel Bonavia, geoffrey Dalfin, Yohan Liaibard, benjamin Thirouard »
 
 Echange sur le raisonnement ecris: "Axel Bonavial"
 
### Autres
- Compréhension de l’exercice et de la distance de hamming + partie écrite : **4h**
- Apprentissage d’un nouveau langage (GOlang) + développement de l’algorithme : **14h**

### Remerciements

Les bienfaits de la [Caféine](https://fr.wikipedia.org/wiki/Caf%C3%A9ine)





