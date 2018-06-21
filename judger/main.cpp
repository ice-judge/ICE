#include <filesystem>
#include <iostream>
namespace fs = std::filesystem;

void copy_foobar() {
  fs::path dir = "/";
  dir /= "sandbox";
  fs::path p = dir / "foobar.txt";
  fs::path copy = p;
  copy += ".bak";
  fs::copy(p, copy);
  fs::path dir_copy = dir;
  dir_copy += ".bak";
  fs::copy(dir, dir_copy, fs::copy_options::recursive);
}

void display_contents(fs::path const & p) {
  std::cout << p.filename() << "\n";

  if (!fs::is_directory(p))
    return;

  for (auto const & e: fs::directory_iterator{p}) {
    if (fs::is_regular_file(e.status())) {
      std::cout << "  " << e.path().filename()
                << " [" << fs::file_size(e) << " bytes]\n";
    } else if (fs::is_directory(e.status())) {
      std::cout << "  " << e.path().filename() << "\n";
    }
  }
}

int main(void)
{
	return 0;
}
