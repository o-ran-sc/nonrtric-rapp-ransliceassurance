from docs_conf.conf import *

#branch configuration

branch = 'f-release'

linkcheck_ignore = [
    'http://localhost.*',
    'http://127.0.0.1.*',
    'https://gerrit.o-ran-sc.org.*',
]

extensions = ['sphinxcontrib.redoc', 'sphinx.ext.intersphinx',]

#intershpinx mapping with other projects
intersphinx_mapping = {}

intersphinx_mapping['nonrtric'] = ('https://docs.o-ran-sc.org/projects/o-ran-sc-nonrtric/en/%s' % branch, None)
