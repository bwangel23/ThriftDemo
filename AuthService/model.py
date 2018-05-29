import sys

from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, String, Integer

engine = create_engine('mysql+mysqldb://root:passwd@localhost:3306/thrift_demo')
Base = declarative_base()


class User(Base):
    __tablename__ = 'user'

    id_ = Column(Integer, primary_key=True)
    username = Column(String(64), nullable=False, index=True)
    password = Column(String(64), nullable=False)
    email = Column(String(64), nullable=False, index=True)

    def __repr__(self):
        return '%s(%r)' % (self.__class__.__name__, self.username)


def usage():
    print("Usage: %s createdb" % sys.argv[0])
    sys.exit(1)


def main():
    if len(sys.argv) != 2:
        usage()

    if sys.argv[1] == "createdb":
        Base.metadata.create_all(engine)
    else:
        usage()


if __name__ == '__main__':
    main()
