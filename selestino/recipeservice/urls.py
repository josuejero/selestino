from django.urls import path
from . import views
import logging
from django.conf import settings  # Add this line
from django.conf.urls.static import static
from django.contrib import admin

logger = logging.getLogger(__name__)

# URLs in recipeservice
def log_request(request, view_name):
    logger.debug(f"Request received for view: {view_name}, path: {request.path}")

urlpatterns = [
    path('', views.recipe_list, name='recipe_list'),  # trailing slash is implicit on the root URL
    path('recipe/<int:id>/', views.recipe_detail, name='recipe_detail'),  # Add trailing slash
    path('recipe/add/', views.add_recipe, name='add_recipe'),  # Add trailing slash
]

if settings.DEBUG:
    urlpatterns += static(settings.STATIC_URL, document_root=settings.STATIC_ROOT)
