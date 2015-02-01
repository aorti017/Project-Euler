var stop = false
for(var i = 1; i <= i; i++){
    if(stop){
        break
    }
    for(var j = 1; j < i; j++){
        if(stop){
            break
        }
        for(var k = 1; k < j; k++){
            if(stop){
                break
            }
            if(((k*k)+(j*j))==(i*i)){
                if(i+j+k == 1000){
                    console.log(i*j*k)
                    stop = true
                }
            }
        }
    }
}
