from PIL import Image as ImagePIL, ImageFont, ImageDraw
import base64
import io
import sys

def main(argv):
    msg="success"
    print("python ", argv)
    return msg

if __name__ == '__main__':
    msg=main(sys.argv[1:])
    print(msg)