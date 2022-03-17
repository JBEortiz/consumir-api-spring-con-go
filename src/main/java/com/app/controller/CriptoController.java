package com.app.controller;

import java.util.List;

import javax.persistence.EntityNotFoundException;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import com.app.entity.Cripto;
import com.app.repository.CriptoRepository;

@RestController
public class CriptoController {

	@Autowired
	private CriptoRepository repository;
	
	@GetMapping("/criptos")
	public ResponseEntity<List<Cripto>> listar() {
		List<Cripto> criptos = repository.findAll();
		if (criptos.isEmpty()) {
			throw new EntityNotFoundException("la lista de criptos esta vacia");
		}
		return ResponseEntity.ok(criptos);
	}

	@GetMapping("/cripto/{id}")
	public ResponseEntity<Cripto> findById(@PathVariable Long id) {
		Cripto cripto;
		cripto = repository.getById(id);
		return ResponseEntity.ok(cripto);

	}
}
