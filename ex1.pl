for(my $i = 0; $i < 1000; $i++){
    if($i % 3 == 0){
        $total += $i
    }
    elsif($i % 5 == 0){
        $total += $i
    }
}
print $total."\n"
