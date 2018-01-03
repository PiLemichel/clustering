# clustering
## Reflexion

## Algorithme

### Détails du Fonctionnement

- L’algorithme crée tout d’abords le maximum de cluster pour notre jeu de donnée. (10 dans notre cas)
- Puis calcule il calcule la distance de Hamming moyenne entre tous ces clusters.
- Compare les distances afin d’extraire la distance minimum entre deux clusters. (si égalité de des distances, il choisit la dernière        distance minimum rencontré) 
- Fusionne les clusters conserver.
- Recommencer depuis l’étape 2 jusqu’à ce que le nombre de cluster généré soit égal au nombre de cluster demandé 

### Exemple (procédure pour 6 clusters) 
 
Cluster de base à 10 elements: 
```
[[{clair 2 2 fine}][{clair 1 2 fine}][{foncée 2 2 fine}] [{clair 2 1 fine}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse}] [{foncée 2 2 épaisse}] [{foncée 1 1 fine}] [{foncée 1 1 épaisse}] [{foncée 2 1 épaisse}]]
```
**{clair 2 2 fine}** et **{clair 1 2 fine}** pour une distance de 1, fusion des deux clusters: 
```
[[{clair 2 2 fine} {clair 1 2 fine}]  [{foncée 2 2 fine}] [{clair 2 1 fine}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse}] [{foncée 2 2 épaisse}] [{foncée 1 1 fine}] [{foncée 1 1 épaisse}] [{foncée 2 1 épaisse}]]
```
**{foncée 2 2 fine}** et **{foncée 2 2 épaisse}** pour une distance de 1, fusion des deux clusters: 
```
[[{clair 2 2 fine} {clair 1 2 fine}]  [{foncée 2 2 fine} {foncée 2 2 épaisse}] [{clair 2 1 fine}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse}] [{foncée 1 1 fine}] [{foncée 1 1 épaisse}] [{foncée 2 1 épaisse}]]
```
**{clair 1 1 épaisse}** et **{foncée 1 1 épaisse}** pour une distance de 1, fusion des deux clusters:  
```
[[{clair 2 2 fine} {clair 1 2 fine}]  [{foncée 2 2 fine} {foncée 2 2 épaisse}] [{clair 2 1 fine}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse} {foncée 1 1 épaisse}]  [{foncée 1 1 fine}] [{foncée 2 1 épaisse}]]
```
**{clair 2 2 fine} {clair 1 2 fine}** et **{clair 2 1 fine}** 3 pour une distance de  1.5, fusion des deux clusters: 
```
 [[{clair 2 2 fine} {clair 1 2 fine} {clair 2 1 fine}] [{foncée 2 2 fine} {foncée 2 2 épaisse}] [{clair 2 2 épaisse}] [{clair 1 1 épaisse} {foncée 1 1 épaisse}]  [{foncée 1 1 fine}] [{foncée 2 1 épaisse}]]
```


