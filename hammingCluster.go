package main

import "fmt"
import "strconv"


type CheckedElements struct {
  element1 int
  element2 int
  distance float64

}

type Element struct {
  couleur string
  noyaux float64
  flagelles float64
  membrane string
}

func main() {
  dataBatch := initialization(); // j'initialise mon jeu de donnée
  nbcluster := 2 // nombre de cluster demandé en sortie
  clustering(nbcluster, dataBatch)
}

func initialization()[]Element{ // initialization function for data batch

  element0 := Element{
    "clair",
    2,
    2,
    "fine",
  }
  element1 := Element{
    "clair",
    1,
    2,
    "fine",
  }
  element2 := Element{
    "foncée",
    2,
    2,
    "fine",
  }
  element3:= Element{
    "clair",
    2,
    1,
    "fine",
  }
  element4:= Element{
    "clair",
    2,
    2,
    "épaisse",
  }
  element5:= Element{
    "clair",
    1,
    1,
    "épaisse",
  }
  element6:= Element{
    "foncée",
    2,
    2,
    "épaisse",
  }
  element7:= Element{
    "foncée",
    1,
    1,
    "fine",
  }
  element8:= Element{
    "foncée",
    1,
    1,
    "épaisse",
  }
  element9:= Element{
    "foncée",
    2,
    1,
    "épaisse",
  }
  //merge des 10 elements
  elementSum := []Element{element0, element1, element2, element3, element4, element5, element6, element7, element8, element9}
  return elementSum
}


// classe et affiche les cluster crée
func clustering( nbCluster int, dataBatch []Element){
  cluster := [10][]Element{} // initialisation de mes clusters
  cpt :=0
  //creation du nombre maximum de clusters
  for cpt < len(dataBatch){
    slice := []Element{dataBatch[cpt]}
    cluster[cpt] = slice
    cpt = cpt+1
  }

  cptCluster := 10 //nombre de cluster maximum pour un jeux de 10 données
  //boucle tant que le nombre de clusters calculer n'est pas egal au nombre de cluster demandé
  for cptCluster != nbCluster{
    // calcule de la distance minimal de hamming entre tous les clusters
    distanceMin := hammingDistanceMin(cluster)
    //fusionner les deux cluster dont la distance est le minimum
    cluster[distanceMin.element1] = append(cluster[distanceMin.element1], cluster[distanceMin.element2]...)
    // supprimer le cluster qui ont etait fusionné
    cluster[distanceMin.element2] = []Element{}
    cptCluster = cptCluster -1
  }
  // affichage des clusters crée
  cptShow := 0
  numCluster:= 0
  //pour chaque cluster dans "cluster"
  for cptShow < len(cluster) {
    //si le cluster n'est pas vide
    if len(cluster[cptShow]) != 0{
      //affiche le cluster
      fmt.Println("Cluster numéro: "+ strconv.Itoa(numCluster))
      fmt.Println(cluster[cptShow])
      numCluster = numCluster + 1
    }
    cptShow = cptShow + 1
  }
}

// génere toute les combinaison possible entre les cluster et renvois la distance moyenne minimum
func hammingDistanceMin(dataBatch [10][]Element)CheckedElements{
  cptElementCheck :=0
  hammingDistances := []CheckedElements{}

  // génere toutes les combinaison possible de distance entre les cluster
  for cptElementCheck < len(dataBatch){
    distance := 0;

    for distance < 5{

      array := []CheckedElements{}
      cpt := 0

      for cpt < len(dataBatch) {
        if len(dataBatch[cptElementCheck]) != 0 || len(dataBatch[cpt])  != 0{
          distanceFind :=  averageDistance(dataBatch[cptElementCheck], dataBatch[cpt])

          if distanceFind != 0{
            checked := CheckedElements{element1: cptElementCheck, element2: cpt, distance: distanceFind}
            array = append(array, checked)
          }
        }
        cpt = cpt + 1
      }
      hammingDistances = append(hammingDistances, array...)
      distance = distance + 1
    }

    cptElementCheck = cptElementCheck + 1
  }

  cptFind :=0
  // initialise une distanceMin Absurde
  minDistance := CheckedElements{
    element1:0,
    element2:0,
    distance:100,
  }
  // trouve la distance de hamming la plus basse dans toutes les distances calculé
  for cptFind < len(hammingDistances){

    if hammingDistances[cptFind].distance < minDistance.distance && hammingDistances[cptFind].element1 != hammingDistances[cptFind].element2{
      minDistance = hammingDistances[cptFind]
    }
    cptFind = cptFind+1
  }

  return minDistance

}
//calcule la distance entre deux cluster
func averageDistance(element1, element2 []Element) float64 {
  cpt1 := 0

  var returnVal float64 = 0
  for cpt1 < len(element1){
    cpt2 := 0
    var diff  float64 = 0
    for cpt2 < len(element2){
      if(element1[cpt1].couleur != element2[cpt2].couleur){
        //check si  la variable "couleur" entre les des deux elements est la même
        diff = diff+1
      }
      if(element1[cpt1].noyaux != element2[cpt2].noyaux){
        //check si  la variable "noyaux" entre les des deux elements est la même
        diff = diff+1
      }
      if(element1[cpt1].flagelles != element2[cpt2].flagelles){
        //check si  la variable "flagelles" entre les des deux elements est la même
        diff = diff+1
      }
      if(element1[cpt1].membrane != element2[cpt2].membrane){
        //check si  la variable "membrane" entre les des deux elements est la même
        diff = diff+1
      }
      cpt2=cpt2 +1
    }
    cpt1 = cpt1+1
    //calcule la distance moyenne entre l'element a l'index cpt1 du cluster "element1" et tous les element du cluster "element2"
    diff = float64(diff)/(float64(cpt2))
    //on addition toutes les moyenne
    returnVal = returnVal + diff
  }
  // on divise la valeur obtenu par le nombre d'element dans le cluster "element1" afin d'obtenir un distance moyenne
  returnVal = float64(returnVal)/float64(cpt1)

  return returnVal
}
