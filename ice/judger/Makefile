CC = g++
FLAGS = -std=c++17 -lstdc++fs
OUTDIR = ./out

all: debug

debug: main.cpp
	$(CC) main.cpp -o $(OUTDIR)/debug-evaluator $(FLAGS)

release: main.cpp
	$(CC) main.cpp -o $(OUTDIR)/evaluator $(FLAGS)

clean:
	rm -rf ./$(OUTDIR)/*


