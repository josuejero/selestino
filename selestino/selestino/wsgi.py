import logging
from django.core.wsgi import get_wsgi_application

logger = logging.getLogger(__name__)

def custom_wsgi_middleware(environ, start_response):
    if 'HTTP_HOST' in environ:
        logger.debug(f"HTTP_HOST: {environ['HTTP_HOST']}")
    return application(environ, start_response)


application = get_wsgi_application()
