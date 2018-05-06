class Lote(object):
	"""docstring for ClassName"""

	def __init__(self, id,owner):
		self.id=id
		self.owner=owner
		self.categoria=''

	def categoria():
		return categoria


class Casa(Lote):
	def __init__(self,id,owner):
		super().__init__(self,id,owner)
		self.categoria='casa'

class Apto(object):
	def __init__(self, arg):
		super(Apto, self).__init__()
		self.arg = arg
		