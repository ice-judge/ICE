GLC = go build
OUTDIR = ./build
RELEASE_FLG = "-s -w"

all: debug

debug: main.go
	$(GLC) -o $(OUTDIR)/debug-scheduler

release: main.go
	$(GLC) -o $(OUTDIR)/scheduler -ldflags $(RELEASE_FLG)

clean:
	rm -rf ./build/*


