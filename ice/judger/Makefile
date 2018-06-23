CC = g++
OUTDIR = ./out
FLAGS = -std=c++17 -lstdc++fs

all: debug

debug: main.cpp
	$(CC) main.cpp -o $(OUTDIR)/debug-judger $(FLAGS)

release: main.cpp
	$(CC) main.cpp -o $(OUTDIR)/judger $(FLAGS)

clean:
	rm -rf ./$(OUTDIR)/*


