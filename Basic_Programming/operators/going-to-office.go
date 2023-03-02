package main
 
import ("fmt")
 
func main(){
    var d int
    var oc int
    var of int
    var od int
    var cs int
    var cb int
    var cm int
    var cd int
    var a int = 0
    var b int = 0
    
    
    fmt.Scan(&d,&oc,&of,&od,&cs,&cb,&cm,&cd)
    
    a=oc+(d-of)*od
    b=cb+(d/cs)*cm+d*cd
    
    if a<=b {
        fmt.Println("Online Taxi\n")
    } else {
        fmt.Println("Classic Taxi\n")
    }
}