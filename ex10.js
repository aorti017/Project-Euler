sum = 0
for(var i = 2; i < 2000000; i++){
    var prime = true
    for(var j = 2; j <= Math.sqrt(i); j++){
        if(j == i){
            continue
        }
        if(i%j==0){
            prime = false
            break
        }
    }
    if(prime){
        sum += i
    }
}
console.log(sum)
