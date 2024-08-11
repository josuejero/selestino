from django.db import models

class Ingredient(models.Model):
    name = models.CharField(max_length=100)
    quantity = models.CharField(max_length=50)

    def __str__(self):
        return self.name

class Recipe(models.Model):
    title = models.CharField(max_length=200)
    description = models.TextField()
    ingredients = models.ManyToManyField(Ingredient)
    instructions = models.TextField()
    prep_time = models.IntegerField(default=0)  # Default to 0 minutes
    cook_time = models.IntegerField(default=0)  # Default to 0 minutes
    total_time = models.IntegerField(editable=False, blank=True, null=True)
    cuisine_type = models.CharField(max_length=100, default='Unknown')  # Add this line
    difficulty = models.CharField(max_length=50, choices=[('Easy', 'Easy'), ('Medium', 'Medium'), ('Hard', 'Hard')], default='Easy')  # Add this line

    def save(self, *args, **kwargs):
        self.total_time = (self.prep_time or 0) + (self.cook_time or 0)
        super().save(*args, **kwargs)

    def __str__(self):
        return self.title

class Review(models.Model):
    recipe = models.ForeignKey(Recipe, on_delete=models.CASCADE, related_name='reviews')
    user = models.CharField(max_length=100)
    rating = models.PositiveSmallIntegerField(choices=[(i, i) for i in range(1, 6)])
    comment = models.TextField(blank=True, null=True)
    created_at = models.DateTimeField(auto_now_add=True)

    def __str__(self):
        return f"{self.user}'s review of {self.recipe.title}"
