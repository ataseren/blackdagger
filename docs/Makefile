# Minimal makefile for Sphinx documentation
#

# You can set these variables from the command line, and also
# from the environment for the first two.
SPHINXOPTS    ?=
SPHINXBUILD   ?= sphinx-build
SOURCEDIR     = source
BUILDDIR      = build

# Put it first so that "make" without argument is like "make help".
help:
	@$(SPHINXBUILD) -M help "$(SOURCEDIR)" "$(BUILDDIR)" $(SPHINXOPTS) $(O)

.PHONY: help Makefile build

# Catch-all target: route all unknown targets to Sphinx using the new
# "make mode" option.  $(O) is meant as a shortcut for $(SPHINXOPTS).
%: Makefile
	@$(SPHINXBUILD) -M $@ "$(SOURCEDIR)" "$(BUILDDIR)" $(SPHINXOPTS) $(O)

local:
	sphinx-autobuild --watch docs --port 8000 --open-browser --delay 1 source docs/build/html

install-dependencies:
	python -m pip install -r requirements.txt

install-libs:
	pip install sphinx sphinx-autobuild

build:
	sphinx-build -b html -D language=us_EN source/ build/html

PYTHON_VER = 3.11.9

venv:
	pyenv local $(PYTHON_VER)
	python -m venv .venv

install-python:
	@pyenv install $(PYTHON_VER)

gen-pot-files:
	sphinx-intl update -p build/gettext -l en