from django.db import models
from django.core.exceptions import ValidationError
from django.core.validators import MinValueValidator, MaxValueValidator, RegexValidator
from django.utils import timezone

class Ingredient(models.Model):
    name = models.CharField(max_length=100)
    quantity = models.CharField(
        max_length=50, 
        validators=[
            RegexValidator(regex=r'^\d+(\.\d{1,2})? [a-zA-Z]+$', message="Enter a valid quantity, e.g., '200 g'.")
        ]
    )

    def __str__(self):
        return self.name

class Recipe(models.Model):
    CUISINE_CHOICES = [
        ('Peruvian', 'Peruvian'),
        ('Italian', 'Italian'),
        ('Japanese', 'Japanese'),
        ('Unknown', 'Unknown'),
    ]

    title = models.CharField(max_length=200, db_index=True)
    description = models.TextField()
    ingredients = models.ManyToManyField(Ingredient)
    instructions = models.TextField()
    prep_time = models.DurationField(default=timezone.timedelta(minutes=0))
    cook_time = models.DurationField(default=timezone.timedelta(minutes=0))
    total_time = models.DurationField(editable=False)
    cuisine_type = models.CharField(max_length=100, choices=CUISINE_CHOICES, default='Unknown')
    difficulty = models.CharField(max_length=50, choices=[('Easy', 'Easy'), ('Medium', 'Medium'), ('Hard', 'Hard')], default='Easy')

    def save(self, *args, **kwargs):
        self.total_time = self.prep_time + self.cook_time
        super().save(*args, **kwargs)

    def __str__(self):
        return self.title

    class Meta:
        ordering = ['title']  
        indexes = [models.Index(fields=['created_at'])]  

class Review(models.Model):
    recipe = models.ForeignKey(Recipe, on_delete=models.CASCADE, related_name='reviews')
    user = models.CharField(max_length=100)
    rating = models.PositiveSmallIntegerField(
        choices=[(i, i) for i in range(1, 6)],
        validators=[MinValueValidator(1), MaxValueValidator(5)]
    )
    comment = models.TextField(blank=True, null=True)
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f"{self.user}'s review of {self.recipe.title} - Rating: {self.rating}"

    class Meta:
        unique_together = ('user', 'recipe')  
