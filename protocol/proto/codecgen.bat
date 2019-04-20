cd ../
codecgen.exe -o codecgen.go login.pb.go game.pb.go coin_flow.go
cd folks
codecgen.exe -o codecgen.go folks.pb.go
cd ../zjh
codecgen.exe -o codecgen.go zjh.pb.go
cd ../fish
codecgen.exe -o codecgen.go fish.pb.go